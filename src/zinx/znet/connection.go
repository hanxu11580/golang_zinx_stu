package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"project_stu/src/zinx/ziface"
)

type Connection struct {
	// 链接当然要有链接
	Conn *net.TCPConn
	// id
	ConnID uint32
	// 状态
	isClosed bool
	// 不同链接有自己不同的处理方法
	// handleAPI ziface.HandFunc
	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
	// 代替原来回调
	Router ziface.IRouter
}

// func NetConnection(conn *net.TCPConn, connID uint32, callback ziface.HandFunc) *Connection {
// 	c := &Connection{
// 		Conn:         conn,
// 		ConnID:       connID,
// 		isClosed:     false,
// 		handleAPI:    callback,
// 		ExitBuffChan: make(chan bool, 1),
// 	}

// 	return c
// }

/*
	原来：传入回调函数 ziface.HandFunc
	现在：传入一个路由对象 ziface.IRouter

	原来：回调函数直接操作连接和数据数组
	现在：路由操作IRequest 并且支持自定义执行阶段前 中 后
*/
func NetConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is  running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	// 未处理粘包问题
	// for {
	// 	//读取我们最大的数据到buf中
	// 	buf := make([]byte, 512)
	// 	_, err := c.Conn.Read(buf)
	// 	if err != nil {
	// 		fmt.Println("recv buf err ", err)
	// 		c.ExitBuffChan <- true
	// 		continue
	// 	}

	// 	// 构造数据
	// 	req := Request{
	// 		Conn: c,
	// 		Data: buf,
	// 	}
	// 	// 操作数据
	// 	go func(request ziface.IRequest) {
	// 		c.Router.PreHandle(request)
	// 		c.Router.Handle(request)
	// 		c.Router.PostHandle(request)
	// 	}(&req)
	// }

	for {
		dp := DataPack{}
		// 8长度的头
		head_data := make([]byte, dp.GetHeadLen())
		// 从这个tcpConnect一直读满这8个字节
		if _, err := io.ReadFull(c.GetTCPConnection(), head_data); err != nil {
			fmt.Println("read msg head error ", err)
			c.ExitBuffChan <- true
			continue
		}
		// 然后解包、获得头部信息
		msg, err := dp.UnPack(head_data)
		if err != nil {
			fmt.Println("unpack error ", err)
			c.ExitBuffChan <- true
			continue
		}

		// 我现在有头部信息了
		var data_bytes []byte
		if msg.GetDataLen() > 0 {
			data_bytes = make([]byte, msg.GetDataLen())
			// 从这个tcpConnect一直读满数据字节
			if _, err := io.ReadFull(c.GetTCPConnection(), data_bytes); err != nil {
				fmt.Println("read msg data error ", err)
				c.ExitBuffChan <- true
				continue
			}
		}

		msg.SetData(data_bytes)

		req := Request{
			Conn: c,
			Msg:  msg,
		}
		go func(request ziface.IRequest) {
			//执行注册的路由方法
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	go c.StartReader()

	for {
		select {
		// 当信道流出 跳出循环
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}

	c.isClosed = true

	c.Conn.Close()
	// 主动关闭
	c.ExitBuffChan <- true
	close(c.ExitBuffChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 提供封包处理
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("Connection closed when send msg")
	}

	dp := NewDataPack()
	msg_bytes, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}

	if _, err := c.Conn.Write(msg_bytes); err != nil {
		fmt.Println("Write msg id ", msgId, " error ")
		c.ExitBuffChan <- true
		return errors.New("conn Write error")
	}

	return nil
}

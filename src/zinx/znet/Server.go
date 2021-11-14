package znet

import (
	"errors"
	"fmt"
	"net"
	"project_stu/src/zinx/utils"
	"project_stu/src/zinx/ziface"
	"time"
)

//iServer 接口实现，定义一个Server服务类
type Server struct {
	//服务器的名称
	Name string
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port int

	MsgHandler ziface.IMsghandler
}

// 链接回调
// 这样是写死的 TODO 我们需要给用户提供自定义行为
// 当前被路由代替
func Callback_Connection(conn *net.TCPConn, data []byte, count int) error {
	fmt.Println("[Conn Handle] execute...")
	if _, err := conn.Write(data[:count]); err != nil {
		fmt.Println("[Conn Handle] write err： ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("开始解析服务IP: %s, Port %d\n", s.IP, s.Port)

	// 这个用于监听
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("解析TCP地址失败: ", err)
			return
		}

		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("监听失败", s.IPVersion, "err", err)
			return
		}

		fmt.Println("开启服务  ", s.Name, " 成功, 监听中...")

		var id_conn uint32
		id_conn = 0

		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept Failed", err)
				continue
			}
			fmt.Println("新连接接入", conn.RemoteAddr().String())
			new_conn := NetConnection(conn, id_conn, s.MsgHandler)
			id_conn++

			// 必须开始新的go程去处理
			go new_conn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
}

func (s *Server) Serve() {
	s.Start()

	for {
		time.Sleep(10 * time.Second)
	}
}

// 服务器添加路由功能
func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	fmt.Println("add router succ")
	s.MsgHandler.AddRouter(msgId, router)
}

func NewServer() ziface.IServer {

	utils.GlobalObject.Reload()

	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
	}

	return s
}

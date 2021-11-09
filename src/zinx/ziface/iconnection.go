package ziface

import "net"

type IConnection interface {
	//启动连接，让当前连接开始工作
	Start()
	//停止连接，结束当前连接状态M
	Stop()

	GetTCPConnection() *net.TCPConn

	GetConnID() uint32

	RemoteAddr() net.Addr

	//直接将Message数据发送数据给远程的TCP客户端
	SendMsg(msgId uint32, data []byte) error
}

// 处理链接
// args: 链接、数据、数量
type HandFunc func(*net.TCPConn, []byte, int) error

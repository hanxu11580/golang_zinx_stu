package znet

import "project_stu/src/zinx/ziface"

type Request struct {
	Conn ziface.IConnection
	// Data []byte
	// 换成Message
	Msg ziface.IMessage
}

func (req *Request) GetConnection() ziface.IConnection {

	return req.Conn
}

func (req *Request) GetData() []byte {

	return req.Msg.GetData()
}

func (req *Request) GetMsgId() uint32 {
	return req.Msg.GetMsgId()
}

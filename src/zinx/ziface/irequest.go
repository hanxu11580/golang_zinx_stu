package ziface

/*
	目的：
		原先的IConnection内回调方法在 Server.go传入
		现在需要用户传入自定义回调
*/

type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetData() []byte            //获取请求消息的数据
	GetMsgId() uint32           //获得消息id
}

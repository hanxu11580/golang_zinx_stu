package ziface

type IMessage interface {
	GetDataLen() uint32 //获取消息数据段长度
	SetDataLen(uint32)  //设置消息数据段长度

	SetMsgId(uint32)  //设置消息ID
	GetMsgId() uint32 //获取消息ID

	GetData() []byte //获取消息内容
	SetData([]byte)  //设置消息内容

}

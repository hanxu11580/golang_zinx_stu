package znet

/*
	Request中只有Data []byte
	没有任何信息就很离谱
	所以加点信息进去
*/

type Message struct {
	Id      uint32 //消息的ID
	DataLen uint32 //消息的长度
	Data    []byte //消息的内容
}

func NewMsgPackage(id uint32, data []byte) *Message {
	msg := Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
	return &msg
}

//获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

//获取消息ID
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

//获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

//设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

//设计消息ID
func (msg *Message) SetMsgId(msgId uint32) {
	msg.Id = msgId
}

//设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}

package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"project_stu/src/zinx/utils"
	"project_stu/src/zinx/ziface"
)

// 数据协议 数据长度-ID-数据

type DataPack struct {
}

//封包拆包实例初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// id uint32 + 数据长度 uint32
	return 8
}

// 封包
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {

	buf := bytes.NewBuffer([]byte{})

	// 数据长度
	if err := binary.Write(buf, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	// id
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	// 数据
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// 解包
// 读出有数据长度和id 但是没有数据的IMessage
func (dp *DataPack) UnPack(data []byte) (ziface.IMessage, error) {
	read_buf := bytes.NewReader(data)
	msg := &Message{}
	// 读取头
	if err := binary.Read(read_buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	// 读msgID
	if err := binary.Read(read_buf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	//
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("Too large msg data recieved")
	}

	return msg, nil
}

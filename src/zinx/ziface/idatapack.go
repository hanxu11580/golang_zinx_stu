package ziface

/*
	针对tcp粘包问题
	次接口提供封包、拆包功能
*/

type IDataPack interface {
	GetHeadLen() uint32                //获取包头长度方法
	Pack(msg IMessage) ([]byte, error) //封包方法
	Unpack([]byte) (IMessage, error)   //拆包方法
}

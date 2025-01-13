package ziface

/*
 针对Message进行TLV格式的封包解包，并得到对应的Message数据包
 TLV格式：
    |   4字节   |    4字节    |      n字节      |
    |   LEN     |    ID      |      DATA       |
 读取数据包读两次，先head后data
*/

// 直接面向TCP链接中的数据流，处理粘包问题
type IDataPack interface {
	GetHeadLen() uint32
	//封包
	Pack(IMessage) ([]byte, error)
	//拆包
	UnPack([]byte) (IMessage, error)
}

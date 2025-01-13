package ziface

/*
 将请求的消息封装到一个message中，定义抽象层接口
*/

type IMessage interface {
	// 获取方法
	GetMsgId() uint32
	GetMsgLen() uint32
	GetData() []byte

	// 设置方法
	SetMsgId(uint32)
	SetMsgLen(uint32)
	SetMsgData([]byte)
}

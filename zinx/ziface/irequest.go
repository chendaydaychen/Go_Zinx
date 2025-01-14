package ziface

// request接口
// 链接信息和数据
type IRequest interface {
	// 获取请求的连接信息
	GetConnection() IConnection
	// 获取请求消息的数据
	GetData() []byte
	// 获取请求的消息ID
	GetMsgID() uint32
}

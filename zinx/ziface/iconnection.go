package ziface

import "net"

// 定义链接模块的抽象层
type IConnection interface {
	//启动链接
	Start()
	//停止链接
	Stop()
	//获取当前连接的socket
	GetTCPConnection() *net.TCPConn
	//获取当前链接ID
	GetConnID() uint32
	//获取远程客户端的TCP状态 IP Port
	RemoteAddr() net.Addr
	//发送数据给远程客户端
	Send(data []byte) error
}

// 定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error

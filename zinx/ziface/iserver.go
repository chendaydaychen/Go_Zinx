package ziface

// 定义服务器接口
type Iserver interface {
	// 启动服务器
	Start()
	// 停止服务器
	Stop()
	// 运行服务器
	Serve()
	// 路由功能，给当前服务器注册一个路由功能，当客户端连接成功后，创建一个对应链接的router
	AddRouter(uint32, IRouter)
	// 获取链接管理
	GetConnMgr() IConnManager
	// 注册OnConnStart 钩子函数的方法
	SetOnConnStart(func(IConnection))
	// 注册OnConnStop 钩子函数的方法
	SetOnConnStop(func(IConnection))
	// 调用OnConnStart 钩子函数的方法
	CallOnConnStart(IConnection)
	// 调用OnConnStop 钩子函数的方法
	CallOnConnStop(IConnection)
}

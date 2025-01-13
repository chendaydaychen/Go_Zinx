package ziface

// 定义服务器接口
type Iserver interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Serve()
	//路由功能，给当前服务器注册一个路由功能，当客户端连接成功后，创建一个对应链接的router
	AddRouter(router IRouter)
}

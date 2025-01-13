package ziface

// IRouter 路由接口
// 路由是请求与处理程序之间的映射,路由里的数据均为IRequest

type IRouter interface {
	// PreHandle 在处理请求之前执行
	PreHandle(request IRequest)
	// Handle 处理请求
	Handle(request IRequest)
	// PostHandle 在处理请求之后执行
	PostHandle(request IRequest)
}

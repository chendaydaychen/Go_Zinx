package ziface

/*
	消息管理抽象层
*/

type IMsgHandle interface {
	//执行对应id的router消息处理方法
	DoMsgHandler(request IRequest)
	//为消息添加具体处理逻辑
	AddRouter(msgId uint32, router IRouter)
}

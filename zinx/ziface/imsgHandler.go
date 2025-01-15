package ziface

/*
	消息管理抽象层
*/

type IMsgHandle interface {
	//执行对应id的router消息处理方法
	DoMsgHandler(IRequest)
	//为消息添加具体处理逻辑
	AddRouter(uint32, IRouter)
	//启动一个worker工作池
	StartWorkerPool()
	//发送消息，根据msgid找到对应的router，并执行router里的业务
	SendMsgToTaskQueue(IRequest)
}

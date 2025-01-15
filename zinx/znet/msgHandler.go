package znet

import (
	"Zinx/zinx/utils"
	"Zinx/zinx/ziface"
	"fmt"
	"strconv"
)

/*
	消息管理模块的实现
*/

type MsgHandle struct {
	// 存放每个MsgID对应的处理方法
	Apis map[uint32]ziface.IRouter
	// 负责Worker取任务的消息队列
	TaskQueue []chan ziface.IRequest
	// 默认的WorkerPool大小
	WorkerPoolSize uint32
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		//从全局配置中获取
		Apis:           make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
	}
}

// 执行对应id的router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	//1. 取出msgID
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgID(), " is not FOUND!")
		return
	}

	//2. 存在，执行对应处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// 为消息添加具体处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	//当前msg绑定的api是否存在
	if _, ok := mh.Apis[msgId]; ok {
		//已经存在
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}
	//添加
	mh.Apis[msgId] = router
	fmt.Println("add api msgId = ", msgId)
}

// 启动一个worker工作池(只发生一次，一个Zinxkerver实例启动一次)
func (mh *MsgHandle) StartWorkerPool() {
	//根据workerPoolSize分别开启worker，每个worker用一个go来承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//给当前worker对应的任务队列开辟空间
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//启动当前worker，阻塞等待channel消息
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

// 启动一个Worker工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started.")
	//不断的等待队列中的消息
	for request := range taskQueue {
		mh.DoMsgHandler(request)
	}
}

// 将消息交给TaskQueue, 由worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//根据ConnID来分配当前的连接应该由哪个worker负责处理
	//轮询的平均分配算法
	//得到需要处理此条连接的workerID
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("workerID = ", workerID, " is working...")
	//将消息发送给worker对应的taskQueue即可
	mh.TaskQueue[workerID] <- request
}

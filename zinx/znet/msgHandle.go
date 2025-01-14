package znet

import (
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
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
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

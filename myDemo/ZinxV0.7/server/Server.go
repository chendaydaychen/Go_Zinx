package main

import (
	"Zinx/zinx/ziface"
	"Zinx/zinx/znet"
	"fmt"
)

/*
	基于Zinx 框架开发的服务器端应用程序
*/

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

// test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle...")
	//先读取客户端数据，再回写ping...ping
	fmt.Println("recv from client: msgId=", request.GetMsgID(),
		", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(200, []byte("ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping...ping err:", err)
	}
}

// hello Zinx test 自定义路由
type HelloZinxRouter struct {
	znet.BaseRouter
}

// test Handle
func (pr *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloRouter Handle...")
	//先读取客户端数据，再回写ping...ping
	fmt.Println("recv from client: msgId=", request.GetMsgID(),
		", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("hello Zinx\n"))
	if err != nil {
		fmt.Println("call back ping...ping err:", err)
	}
}

func main() {
	//创建一个server 服务器
	server := znet.NewServer("zinx-server v0.7")

	//给当前zinx框架添加一个自定义的router
	server.AddRouter(0, &PingRouter{})
	server.AddRouter(1, &HelloZinxRouter{})
	//启动server 服务器
	server.Serve()
}

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

// test PreHandle
func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	if err != nil {
		fmt.Println("call back before ping err:", err)
	}

}

// test Handle
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping...ping err:", err)
	}

}

// test PostHandle
func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	if err != nil {
		fmt.Println("call back after ping err:", err)
	}
}

func main() {
	//创建一个server 服务器
	server := znet.NewServer("zinx-server v0.4")

	//给当前zinx框架添加一个自定义的router
	server.AddRouter(&PingRouter{})
	//启动server 服务器
	server.Serve()
}

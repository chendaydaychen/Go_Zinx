package main

import (
	"Zinx/zinx/znet"
)

/*
	基于Zinx 框架开发的服务器端应用程序
*/

func main() {
	//创建一个server 服务器
	server := znet.NewServer("zinx-server v0.1")
	//启动server 服务器
	server.Serve()
}

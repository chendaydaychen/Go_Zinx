package main

import "Zinx/zinx/znet"

func main() {
	// 创建server句柄
	s := znet.NewServer("[MMO Game Zinx]")

	// 链接创建销毁的Hook函数

	// 注册一些路由服务

	// 启动服务
	s.Serve()
}

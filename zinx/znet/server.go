package znet

import "Zinx/zinx/ziface"

// iServer接口实现 定义一个Server的服务
type Server struct {
	Name    string //服务器的名称
	Version string //服务器版本
	IP      string //服务器监听的地址
	Port    int    //服务器监听端口
}

// 启动server的服务功能
func (s *Server) Start() {

}

// 停止server的服务功能
func (s *Server) Stop() {

}

// 启动server服务
func (s *Server) Serve() {

}

// 初始化server的方法
func NewServer(name string) ziface.IServer {
	return &Server{
		Name:    name,
		Version: "tcp4",
		IP:      "0.0.0.0",
		Port:    8999,
	}
}

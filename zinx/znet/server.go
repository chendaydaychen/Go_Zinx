package znet

import (
	"Zinx/zinx/ziface"
	"fmt"
	"net"
)

// iServer接口实现 定义一个Server的服务
type Server struct {
	Name    string //服务器的名称
	Version string //服务器版本
	IP      string //服务器监听的地址
	Port    int    //服务器监听端口
}

// 启动server的服务功能
func (s *Server) Start() {
	go func() {
		// 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.Version, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("net.ResolveTCPAddr err:", err)
			return
		}

		// 监听
		listener, err := net.ListenTCP(s.Version, addr)
		if err != nil {
			fmt.Println("net.ListenTCP err:", err)
			return
		}
		defer listener.Close()
		fmt.Println("服务器启动成功，监听地址：", addr.String())
		// 阻塞等待链接，处理客户端业务
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listener.AcceptTCP err:", err)
				continue
			}
			//已经建立连接，处理业务，做一个最基本最大512字节长度的回显业务
			go func() {
				for {
					buf := make([]byte, 512)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println("conn.Read err:", err)
						break
					}
					if n == 0 {
						fmt.Println("客户端退出")
						break
					}
					fmt.Println("收到客户端的数据：", string(buf))
					conn.Write(buf[:n])
				}
			}()
		}
	}()
}

// 停止server的服务功能
func (s *Server) Stop() {
	// TODO 停止server，将资源状态和一些连接信息停止接收或回收
}

// 启动server服务
func (s *Server) Serve() {
	// 启动server
	s.Start()

	// TODO 做额外业务

	// 阻塞状态
	select {}
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

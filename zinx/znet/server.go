package znet

import (
	"Zinx/zinx/ziface"
	"errors"
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

// 定义当前客户端所绑定的API，目前写死，以后优化
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	//回显
	fmt.Println("[Zinx] CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err:", err)
		return errors.New("write back buf err")
	}
	return nil
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

		var cid uint32 = 0
		defer listener.Close()
		fmt.Println("服务器启动成功，监听地址：", addr.String())
		// 阻塞等待链接，处理客户端业务
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listener.AcceptTCP err:", err)
				continue
			}
			//处理新连接的业务方法和conn绑定
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++

			//启动
			dealConn.Start()
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
func NewServer(name string) ziface.Iserver {
	return &Server{
		Name:    name,
		Version: "tcp4",
		IP:      "0.0.0.0",
		Port:    8999,
	}
}

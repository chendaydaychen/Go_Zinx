package znet

import (
	"Zinx/zinx/utils"
	"Zinx/zinx/ziface"
	"fmt"
	"net"
)

// iServer接口实现 定义一个Server的服务
type Server struct {
	Name       string            // 服务器的名称
	Version    string            // 服务器版本
	IP         string            // 服务器监听的地址
	Port       int               // 服务器监听端口
	MsgHandler ziface.IMsgHandle // 当前server的消息管理模块，用来绑定msgid对应的处理业务Api关系
}

// 启动server的服务功能
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name : %s, Version: %s, IP: %s, Port: %d",
		utils.GlobalObject.Name,
		utils.GlobalObject.Version,
		utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort)
	fmt.Println("服务器开始启动...")

	go func() {
		// 0 开始消息队列和worker pool
		s.MsgHandler.StartWorkerPool()
		// 1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.Version, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("net.ResolveTCPAddr err:", err)
			return
		}

		// 2 监听
		listener, err := net.ListenTCP(s.Version, addr)
		if err != nil {
			fmt.Println("net.ListenTCP err:", err)
			return
		}

		var cid uint32 = 0
		defer listener.Close()
		fmt.Println("服务器启动成功，监听地址：", addr.String())
		// 3 阻塞等待链接，处理客户端业务
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("listener.AcceptTCP err:", err)
				continue
			}
			// 处理新连接的业务方法和conn绑定
			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++

			// 启动
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

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	// 添加路由
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("router add success")
}

// 初始化server的方法
func NewServer(name string) ziface.Iserver {
	return &Server{
		Name:       utils.GlobalObject.Name,
		Version:    "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
	}
}

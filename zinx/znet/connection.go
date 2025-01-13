package znet

import (
	"Zinx/zinx/utils"
	"Zinx/zinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	ExitChan chan bool
	Router   ziface.IRouter
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,               // 链接
		ConnID:   connID,             // 连接ID
		isClosed: false,              // 链接是否关闭
		Router:   router,             // 路由
		ExitChan: make(chan bool, 1), // 退出消息
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("connID=", c.ConnID, " reader is running")
	defer fmt.Println("connID=", c.ConnID, " reader exit")

	defer c.Stop()

	for {
		//读数据到buf,最大大数据为MaxPackageSize
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read from client err:", err)
			continue
		}
		// 得到当前conn数据的Request请求数据
		req := &Request{
			conn: c,
			data: buf,
		}
		// 执行注册路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(req)

		// 将buf数据传递给router，调用路由方法，从路由中找到注册绑定的的Conn对应router

	}
}

// 启动链接 读写分离
func (c *Connection) Start() {
	fmt.Println("connID=", c.ConnID, " start")

	// 启动从当前读数据的业务
	go c.StartReader()
	//TODO 启动从当前写数据的业务
}

// 停止链接
func (c *Connection) Stop() {
	fmt.Println("connID=", c.ConnID, " stop")

	//如果连接已经关闭
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	//回收资源
	close(c.ExitChan)
}

// 获取当前连接的socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态 IP Port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据给远程客户端
func (c *Connection) Send(data []byte) error {
	return nil
}

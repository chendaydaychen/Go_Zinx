package znet

import (
	"Zinx/zinx/ziface"
	"fmt"
	"net"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	isClosed  bool
	handleAPI ziface.HandleFunc
	ExitChan  chan bool
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callbackAPI,
		ExitChan:  make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("connID=", c.ConnID, " reader is running")
	defer fmt.Println("connID=", c.ConnID, " reader exit")

	defer c.Stop()

	for {
		//读数据到buf,最大512字节
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read from client err:", err)
			continue
		}

		//调用当前所绑定的API业务
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("call back err:", err)
			break
		}

	}
}

// 启动链接
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

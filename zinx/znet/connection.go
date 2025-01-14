package znet

import (
	"Zinx/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	Conn       *net.TCPConn
	ConnID     uint32
	isClosed   bool
	ExitChan   chan bool
	MsgHandler ziface.IMsgHandle
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msghandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:       conn,               // 链接
		ConnID:     connID,             // 连接ID
		isClosed:   false,              // 链接是否关闭
		MsgHandler: msghandler,         // 路由
		ExitChan:   make(chan bool, 1), // 退出消息
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("connID=", c.ConnID, " reader is running")
	defer fmt.Println("connID=", c.ConnID, " reader exit")

	defer c.Stop()

	for {
		//读数据到buf,最大大数据为MaxPackageSize
		// buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		// _, err := c.Conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("read from client err:", err)
		// 	continue
		// }
		// 创建一个拆包解包对象
		dp := NewDataPack()

		//读取head，八字节
		headData := make([]byte, dp.GetHeadLen())
		_, err := c.Conn.Read(headData)
		if err != nil {
			fmt.Println("read head err:", err)
			break
		}

		//拆包 得到msgId和msgLen，放到msg中
		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack err:", err)
			break
		}

		var data []byte
		//根据datalen，再次读取data，放到msg中
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read data err:", err)
				break
			}
		}
		msg.SetMsgData(data)

		// 得到当前conn数据的Request请求数据
		req := &Request{
			conn: c,
			msg:  msg,
		}
		// 将buf数据传递给router，调用路由方法，从路由中找到注册绑定的的Conn对应router
		go c.MsgHandler.DoMsgHandler(req)
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

// 提供一个SendMessage方法，将我们要发送给客户端的数据先封包再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection closed when send msg")
	}

	//将data进行封包
	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("pack error msg id = ", msgId)
		return errors.New("pack error msg")
	}
	//将data发送给客户端
	if _, err := c.Conn.Write(binaryMsg); err != nil {
		fmt.Println("write back buf err:", err)
		return errors.New("conn Write error")
	}

	return nil
}

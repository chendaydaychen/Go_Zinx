package znet

import (
	"Zinx/zinx/utils"
	"Zinx/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	Conn       *net.TCPConn      // 链接
	ConnID     uint32            // 连接ID
	isClosed   bool              // 链接是否关闭
	ExitChan   chan bool         // 退出信号(由Reader告知Writer)
	msgChan    chan []byte       // 无缓冲消息队列
	MsgHandler ziface.IMsgHandle // 消息处理模块
}

// 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msghandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		MsgHandler: msghandler,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
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
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
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
		if utils.GlobalObject.WorkerPoolSize > 0 {
			//已经启动工作池机制，将消息交给Worker处理
			c.MsgHandler.SendMsgToTaskQueue(req)
		} else {
			//从路由中找到注册绑定的Conn对应的router，然后执行router的Handle方法
			go c.MsgHandler.DoMsgHandler(req)
		}
	}
}

// 专门发给客户端消息的模块
func (c *Connection) StratWriter() {

	fmt.Println("connID=", c.ConnID, " writer is running")
	defer fmt.Println("connID=", c.ConnID, " writer exit")
	//主要监控一个管道,不断阻塞等待，回写给用户
	for {
		select {
		case <-c.ExitChan:
			//如果当前链接已经关闭
			return
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("SendMsg err:", err)
				return
			}
		}
	}
}

// 启动链接 读写分离
func (c *Connection) Start() {
	fmt.Println("connID=", c.ConnID, " start")

	// 启动从当前读数据的业务
	go c.StartReader()
	//TODO 启动从当前写数据的业务
	go c.StratWriter()
}

// 停止链接
func (c *Connection) Stop() {
	fmt.Println("connID=", c.ConnID, " stop")
	//如果连接已经关闭
	if c.isClosed {
		return
	}
	c.isClosed = true
	//关闭socket链接
	c.Conn.Close()
	//告知Writer退出
	c.ExitChan <- true
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
	// //将data发送给客户端
	// if _, err := c.Conn.Write(binaryMsg); err != nil {
	// 	fmt.Println("write back buf err:", err)
	// 	return errors.New("conn Write error")
	// }
	//将data发送给msgChan
	c.msgChan <- binaryMsg
	return nil
}

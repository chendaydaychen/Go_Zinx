package main

import (
	"Zinx/zinx/znet"
	"fmt"
	"io"
	"net"
	"time"
)

// 创建客户端实例
func main() {
	fmt.Println("client start")

	time.Sleep(time.Second)

	// 创建
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	defer conn.Close()

	//链接调用Write写数据
	for {
		//发送封包的message消息
		dp := znet.NewDataPack()
		binaryMsg, _ := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx V0.7 client test0 message")))
		_, err := conn.Write(binaryMsg)
		if err != nil {
			fmt.Println("client write err:", err)
			return
		}

		// _, err := conn.Write([]byte("hello zinx v0.1"))
		// if err != nil {
		// 	fmt.Println("write conn err:", err)
		// 	return
		// }
		// buf := make([]byte, 512)
		// cnt, err := conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("read buf err:", err)
		// 	return
		// }
		// fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)

		//服务端回复数据，拆包
		msgHeadData := make([]byte, dp.GetHeadLen())
		//先读取head部分，得到id和datalen
		io.ReadFull(conn, msgHeadData)
		msg, err := dp.UnPack(msgHeadData)
		if err != nil {
			fmt.Println("unpack err:", err)
			return
		}
		//再根据len读data部分
		if msg.GetMsgLen() > 0 {
			msgData := make([]byte, msg.GetMsgLen())
			io.ReadFull(conn, msgData)
			msg.SetMsgData(msgData)
			fmt.Println("--->Recv Server Msg: ID=", msg.GetMsgId(), ", len=", msg.GetMsgLen(), ", data=", string(msg.GetData()))
		}

		//cpu阻塞
		time.Sleep(time.Second)
	}

}

package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

// 测试数据包封包拆包
func TestDataPack(t *testing.T) {
	/*
		模拟的服务器
	*/
	// 1.创建一个server句柄，使用net.Listen("tcp", "127.0.0.1:7777")
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		t.Fatal(err)
	}

	// 启动一个go，承载从客户端处理业务
	go func() {
		// 2.从客户端读数据，拆包处理
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("listener accept err:", err)
			}
			go func(conn net.Conn) {
				// 处理客户端请求
				// 创建一个dp
				dp := NewDataPack()
				for {
					// 第一次读出head
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error")
						break
					}

					// 将headData字节流进行拆包，得到msgId和dataLen
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("unpack err")
						return
					}
					if msgHead.GetMsgLen() > 0 {
						// 有dataLen
						// 第二次读出data
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgLen())
						// 根据dataLen进行第二次读取
						_, err := io.ReadFull(conn, msg.GetData())
						if err != nil {
							fmt.Println("read msg data error")
							return
						}

						fmt.Println("msgId: ", msg.GetMsgId(), "msgLen: ", msg.GetMsgLen(), "msgData: ", string(msg.GetData()))
					}

				}
			}(conn)
		}
	}()

	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("dial err:", err)
		return
	}
	defer conn.Close()

	// 创建一个封包对象
	dp := NewDataPack()
	// 创建一个msg
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	// 进行封包
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("msg1 pack err")
		return
	}
	msg2 := &Message{
		Id:      1,
		DataLen: 9,
		Data:    []byte{'a', 'h', 'n', 'x', 'o', 't', 'e', 's', 't'},
	}
	// 进行封包
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("msg2 pack err")
		return
	}

	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)

	select {}
}

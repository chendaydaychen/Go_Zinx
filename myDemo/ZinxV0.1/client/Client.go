package main

import (
	"fmt"
	"net"
	"time"
)

// 创建客户端实例
func main() {
	fmt.Println("client start")

	time.Sleep(time.Second)

	// 创建客户端实例
	conn, err := net.Dial("tcp", "0.0.0.0:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}
	defer conn.Close()
	//链接调用Write写数据
	for {
		_, err := conn.Write([]byte("hello zinx v0.1"))
		if err != nil {
			fmt.Println("write conn err:", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:", err)
			return
		}
		fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)

		//cpu阻塞
		time.Sleep(time.Second)
	}

}

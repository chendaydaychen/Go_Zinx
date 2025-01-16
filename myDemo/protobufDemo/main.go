package main

import (
	"Zinx/myDemo/protobufDemo/pb"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	//定义一个Person结构对象
	person := &pb.Person{
		Name: "小明",
		Age:  18,
		Emails: []string{
			"123@qq.com",
			"456@qq.com",
		},
		Phones: []*pb.PhoneNumber{
			{
				Number: "123456789",
				Type:   pb.PhoneType_MOBILE,
			},
			{
				Number: "987654321",
				Type:   pb.PhoneType_HOME,
			},
		},
	}

	//将person对象数据序列化 即将protobuf的message进行序列化，得到二进制文件
	data, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	fmt.Println(data)

	newdata := &pb.Person{}
	err = proto.Unmarshal(data, newdata)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	fmt.Println(newdata)

}

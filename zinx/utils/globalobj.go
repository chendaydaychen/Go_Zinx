package utils

import (
	"Zinx/zinx/ziface"
	"encoding/json"
	"os"
)

/*
	存储一切有关Zinx框架的全局参数，供其他模块使用
	一些参数可以有zinx.json获取
*/

type GlobalObj struct {
	//Server
	TcpServer ziface.Iserver // 当前Zinx框架的全局Server对象
	Host      string         // 当前服务器主机IP
	TcpPort   int            // 当前服务器主机监听端口号
	Name      string         // 当前服务器名称

	//Zinx

	Version        string // 当前Zinx的版本号
	MaxConn        int    // 当前服务器主机允许的最大链接个数
	MaxPackageSize uint32 // 当前Zinx框架数据包最大值
}

/*
	定义一个全局的对外GlobalObj
*/

var GlobalObject *GlobalObj

/*
	从zinx.json读取参数
*/

func (g *GlobalObj) Reload() {
	// 从配置文件里面加载一些参数
	data, err := os.ReadFile("../conf/zinx.json")
	if err != nil {
		panic(err)
	}
	// 将json文件解析到全局变量
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/*
	提供init方法，初始化当前GlobalObject
*/

func init() {
	// 初始化GlobalObject变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "v0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	GlobalObject.Reload()
}

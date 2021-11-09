package utils

import (
	"encoding/json"
	"io/ioutil"
	"project_stu/src/zinx/ziface"
)

type GlobalObj struct {
	TcpServer ziface.IServer //当前Zinx的全局Server对象
	Host      string         //当前服务器主机IP
	TcpPort   int            //当前服务器主机监听端口号
	Name      string         //当前服务器名称
	Version   string         //当前Zinx版本号

	MaxPacketSize uint32 // 数据包的最大值
	MaxConn       int    // 当前服务器主机允许的最大链接个数
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {

	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	if GlobalObject == nil {
		GlobalObject = new(GlobalObj)
	}

	err = json.Unmarshal(data, GlobalObject)
	if err != nil {
		panic(err)
	}
}

//
func Init() {

	GlobalObject = &GlobalObj{
		Host:          "127.0.0.1",
		TcpPort:       7777,
		Name:          "[Zinx Tcp Server]",
		Version:       "[v_]",
		MaxPacketSize: 4096, // 8kb
		MaxConn:       12000,
	}

	GlobalObject.Reload()
}

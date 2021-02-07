/**
 *@Desc:存储一切有关Zinx框架的全局参数，供其他模块使用
 *@Author:Giousa
 *@Date:2021/1/6
 */
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinx/ziface"
)

type GlobalObj struct {

	TcpServer ziface.IServer

	Host string

	TcpPort int

	Name string

	Version string

	//允许最大链接数
	MaxConn int

	//包的最大值
	MaxPackageSize uint32
}


var GlobalObject *GlobalObj

func (g *GlobalObj) Reload()  {

	//TODO 后期针对这个，做相对文件夹访问处理
	//data,err := ioutil.ReadFile("/Users/zhangmengmeng/Documents/CodeResource/go_project/src/zinxProject/zinxV0.4/conf/zinx.json")
	data,err := ioutil.ReadFile("conf/zinx.json")
	if err != nil{
		fmt.Println("GlobalObj file Read err:",err)
		return
		//panic(err)
	}
	//将json文件数据解析到struct中
	err = json.Unmarshal(data,&GlobalObject)
	if err != nil{
		fmt.Println("GlobalObj file Unmarshal err:",err)
		//panic(err)
		return
	}
}

//初始化方法
func init()  {
	fmt.Println("GlobalObj init...")
	GlobalObject = &GlobalObj{
		Name: "ZinxServerAPP",
		Version: "V0.4",
		TcpPort: 999,
		Host: "0.0.0.0",
		MaxConn: 1000,
		MaxPackageSize: 4096,
	}

	//从conf/zinx.json
	//GlobalObject.Reload()
}
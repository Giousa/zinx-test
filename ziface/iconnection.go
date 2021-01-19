/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/4
 */
package ziface

import "net"

type IConnection interface {
	//启动链接Start()
	Start()

	//停止链接Stop()
	Stop()

	//获取当前链接的conn对象（套接字）
	GetTCPConnection() *net.TCPConn

	//得到链接ID
	GetConnID() uint32

	//得到客户端链接的地址和端口
	RemoteAddr() net.Addr

	//发送数据的方法Send()
	SendMsg(msgId uint32,data []byte) error
}


//定义一个处理链接业务的方法
type HandlerFunc func(*net.TCPConn,[]byte,int) error
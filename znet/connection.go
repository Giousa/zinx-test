/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/4
 */
package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	//socket TCP套接字
	Conn *net.TCPConn
	//链接的ID
	ConnID uint32
	//当前链接的状态（是否已关闭）
	isClosed bool
	//与当前链接所绑定的业务处理方法
	//HandlerAPI ziface.HandlerFunc
	//告知当前链接已经退出/停止的channel
	ExitChan chan bool
	//该链接，处理的方法Router
	Router ziface.IRouter
}

//初始化链接模块
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {

	c := Connection{
		Conn: conn,
		ConnID: connID,
		Router: router,
		isClosed: false,
		ExitChan: make(chan bool,1),
	}

	return &c
}

func (c *Connection) StartReader()  {
	fmt.Println("Reader Goroutine is Running...")

	defer fmt.Println("ConnId = ",c.ConnID," Reader is exit,remote addr is ",c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte,1024)
		//n,err := c.Conn.Read(buf)
		//if err != nil{
		//	fmt.Println("reader err : ",err)
		//	break
		//}
		_,err := c.Conn.Read(buf)
		if err != nil{
			fmt.Println("reader err : ",err)
			break
		}

		req := Request{
			conn: c,
			data: buf,
		}

		//执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		//我们从router里面调用
		////调用当前链接所绑定的HandlerAPI
		//if err := c.HandlerAPI(c.Conn,buf,n);err != nil{
		//	fmt.Println("ConnID ",c.ConnID," handler is err : ",err)
		//	break
		//}
	}

}

//启动链接Start()
func (c *Connection) Start(){
	fmt.Println("Conn start...ConnID = ",c.ConnID)
	//启动从当前链接的读数据业务
	go c.StartReader()
	//TODO 启动写的数据业务
}

//停止链接Stop()
func (c *Connection) Stop(){
	fmt.Println("Conn stop...ConnID = ",c.ConnID)

	if c.isClosed{
		return
	}

	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

//获取当前链接的conn对象（套接字）
func (c *Connection) GetTCPConnection() *net.TCPConn{

	return c.Conn
}

//得到链接ID
func (c *Connection) GetConnID() uint32{

	return c.ConnID
}

//得到客户端链接的地址和端口
func (c *Connection) RemoteAddr() net.Addr{

	return c.Conn.RemoteAddr()
}

//发送数据的方法Send()
func (c *Connection) Send(data []byte) error{

	return nil
}
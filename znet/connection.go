/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/4
 */
package znet

import (
	"errors"
	"fmt"
	"io"
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
	//Router ziface.IRouter
	MsgHandler ziface.IMsgHandler
}

//初始化链接模块
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {

	c := Connection{
		Conn: conn,
		ConnID: connID,
		MsgHandler: msgHandler,
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
		//buf := make([]byte,1024)
		//
		//_,err := c.Conn.Read(buf)
		//if err != nil{
		//	fmt.Println("reader err : ",err)
		//	break
		//}

		//拆包
		dp := NewDataPack()

		headData := make([]byte,dp.GetHeadLen())
		_,err := io.ReadFull(c.GetTCPConnection(),headData)
		if err != nil{
			fmt.Println("conn read first err ",err)
			break
		}

		msg,err := dp.Unpack(headData)
		if err != nil{
			fmt.Println("unpack err ",err)
			return
		}

		dataLen := msg.GetDataLen()

		var data []byte
		if dataLen > 0{
			data = make([]byte,dataLen)
			_,err := io.ReadFull(c.GetTCPConnection(),data)
			if err != nil{
				fmt.Println("conn read second err ",err)
				break
			}

		}

		msg.SetData(data)

		req := Request{
			conn: c,
			msg: msg,
		}

		//执行注册的路由方法
		//go func(request ziface.IRequest) {
		//	c.Router.PreHandle(request)
		//	c.Router.Handle(request)
		//	c.Router.PostHandle(request)
		//}(&req)
		go c.MsgHandler.DoMsgHandler(&req)

		//我们从router里面调用
		////调用当前链接所绑定的HandlerAPI
		//if err := c.HandlerAPI(c.Conn,buf,n);err != nil{
		//	fmt.Println("ConnID ",c.ConnID," handler is err : ",err)
		//	break
		//}
	}

}

//服务端发送数据,需要将数据，封包后发送
func (c *Connection) SendMsg(msgId uint32,data []byte) error{

	if c.isClosed{
		return errors.New("Connection has closed")
	}

	//将数据进行封包
	dp := NewDataPack()
	msg := NewMsgPackage(msgId,data)

	binaryMsg,err := dp.Pack(msg)
	if err != nil{
		fmt.Println("Pack msg error ",err)
		return errors.New("Pack msg error")
	}

	_,errW := c.GetTCPConnection().Write(binaryMsg)
	if errW != nil {
		fmt.Println("Write msg error ",err)
		return errors.New("Write msg error")
	}

	return nil
}


//启动链接Start()
func (c *Connection) Start(){
	fmt.Println("Conn start...ConnID = ",c.ConnID)
	//启动从当前链接的读数据业务
	go c.StartReader()
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

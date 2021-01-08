/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/3
 */
package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

type Server struct {

	Name string
	IPVersion string
	IP string
	Port int
	Router ziface.IRouter

}

//回调函数
//func CallBackToClient(conn *net.TCPConn,data []byte,n int) error{
//	fmt.Println("[Conn Handler Callback To Client...] ： ",string(data[:n]))
//	if _,err := conn.Write(data[:n]) ; err != nil{
//		fmt.Println("write back buf err ",err)
//		return errors.New("CallBackToClient error")
//	}
//	return nil
//}

func (s *Server) Start()  {

	fmt.Println("【Server】",utils.GlobalObject.Name," is Starting...")
	fmt.Printf("[Start]Server listen at IP:%s,Port:%d,is Starting...\n",s.IP,s.Port)

	go func() {
		addr ,err := net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
		if err != nil{
			fmt.Println("resoleveTcp err",err)
			return
		}
		listern,err := net.ListenTCP(s.IPVersion,addr)
		if err != nil{
			fmt.Println("ListenTCP err",err)
			return
		}

		defer listern.Close()

		fmt.Printf("Name:%s Server Start Success!\n",s.Name)

		var cid uint32
		cid = 0
		for{
			conn,err := listern.AcceptTCP()
			clientAddr := conn.RemoteAddr().String()
			fmt.Printf("[%v] 连接成功\n",clientAddr)

			if err != nil{
				fmt.Println("Accept err",err)
				continue
			}

			dealConn := NewConnection(conn,cid,s.Router)
			cid++


			//启动当前的链接业务处理
			go dealConn.Start()

			//go func() {
			//	buf := make([]byte,1024)
			//	for{
			//		n,err := conn.Read(buf)
			//		if n == 0{
			//			return
			//		}
			//
			//		if err != nil{
			//			fmt.Println("Read err",err)
			//			continue
			//		}
			//
			//		if _,err := conn.Write(buf[:n]);err != nil{
			//			fmt.Println("writer back buf err ",err)
			//			continue
			//		}
			//
			//		fmt.Println("Server Read ",string(buf[:n]))
			//	}
			//}()

		}
	}()

}

func (s *Server) Stop(){

}

func (s *Server) Run() {
	//启动server服务功能
	s.Start()

	//TODO 启动服务器之后的额外业务

	//阻塞状态 必须的，不然Start里面的go会直接略过
	select {

	}

}

func (s *Server) AddRouter(router ziface.IRouter)  {
	s.Router = router
	fmt.Println("Add Router Success!!")
}

func NewServer(name string) ziface.IServer {
	fmt.Println("NewServer 开始创建")
	s := &Server{
		Name: utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP: utils.GlobalObject.Host,
		Port: utils.GlobalObject.TcpPort,
		Router: nil,
	}
	return s
}
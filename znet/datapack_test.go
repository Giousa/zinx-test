/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/19
 */
package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	//模拟服务器
	//创建SocketTCP
	fmt.Println("准备开始模拟服务端：")
	listener,err := net.Listen("tcp","127.0.0.1:7777")
	if err != nil{
		fmt.Println("server listen err ",err)
		return
	}

	//创建一个go，复制从客户端处理业务
	go func() {
		fmt.Println("等待客户端连接......")
		for{
			conn,err := listener.Accept()
			if err != nil{
				fmt.Println("server accept err ",err)
				return
			}

			fmt.Println("客户端连接成功！")
			go func(conn net.Conn) {
				//从客户端读取数据，拆包处理
				//------拆包------
				//第一次从conn读，读取head
				//第二次从conn读，根据head的dataLen，读取内容

				dp := NewDataPack()
				for{

					//读取长度是8，仅仅读取了头部
					headData := make([]byte,dp.GetHeadLen())


					//第一次读取
					n,err := io.ReadFull(conn,headData)
					if err != nil{
						fmt.Println("server read head err ",err)
						break
					}
					//读取长度： 8 这个是head头部数据长度
					//接下来，我们需要解析head数据，获取数据body长度，根据这个长度进行读取
					fmt.Println("读取长度：",n)
					//[4 0 0 0 1 0 0 0 122 105 110 120
					//15 0 0 0 2 0 0 0 228 189 160 229 165 189 229 149 138 229 133 136 231 148 159]
					//[4 0 0 0 1 0 0 0]
					fmt.Println("服务端读取头部信息：",headData)

					//解包，获取头部信息
					msgHead,err := dp.Unpack(headData)

					if err != nil{
						fmt.Println("server unpack err ",err)
						return
					}

					fmt.Println("服务端读取头部：",msgHead)

					dataLen := msgHead.GetDataLen()

					fmt.Println("服务端从头部读取dataLen：",dataLen)

					if dataLen > 0{
						//msg是有数据的，需要进行第二次读取

						//转换类型
						msg := msgHead.(*Message)

						//data  := make([]byte,dataLen)
						//_,err := io.ReadFull(conn,data)
						msg.Data  = make([]byte,dataLen)
						_,err := io.ReadFull(conn,msg.Data)


						if err != nil{
							fmt.Println("server unpdack data err ",err)
							return
						}
						fmt.Println("服务端读取data：",msg.Data)
						fmt.Println("服务端读取data：",string(msg.Data))
						fmt.Println("------读取完毕------")
						//fmt.Printf("id = %v,len = %v,data = %v \n",msg.Id,msg.DataLen,string(msg.Data))

					}

				}


			}(conn)
		}
	}()




	fmt.Println("准备开始模拟客户端：")
	//模拟客户端
	conn,err := net.Dial("tcp","127.0.0.1:7777")

	if err != nil{
		fmt.Println("client dial err ",err)
		return
	}

	//创建一个封包对象
	dp := NewDataPack()
	//模拟粘包，封装2个msg一起发生
	//封装第一个
	bys1 := []byte{'z','i','n','x'}
	msg1 := &Message{
		Id: 1,
		DataLen: uint32(len(bys1)),
		Data: bys1,
	}
	sendData1,err := dp.Pack(msg1)
	if err != nil{
		fmt.Println("client pack msg1 err ",err)
		return
	}

	//封装第二个
	bys2 := []byte("你好啊先生")
	msg2:= &Message{
		Id: 2,
		DataLen: uint32(len(bys2)),
		Data: bys2,
	}
	sendData2,err := dp.Pack(msg2)
	if err != nil{
		fmt.Println("client pack msg2 err ",err)
		return
	}

	//将两个包黏在一起
	//...打散
	sendData1 = append(sendData1,sendData2...)

	fmt.Println("客户端发送数据str：",string(sendData1))
	fmt.Println("客户端发送数据[]byte：",sendData1)
	//发送
	conn.Write(sendData1)

	//客户端阻塞
	select {

	}

}

/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/14
 */
package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"zinx/utils"
	"zinx/ziface"
)

//封包、拆包的具体模块
type DataPack struct {

}

//拆包封包实例的初始方法
func NewDataPack() *DataPack {
	return &DataPack{}
}


//包的结构：
// head(DataLen id)  body(Data)
//type Message struct {
//	Id      uint32 //消息id
//	DataLen uint32 //消息长度
//	Data    []byte //消息内容
//}

//获取包的头的长度方法
func (d *DataPack) GetHeadLen() uint32{
	//DataLen uint32 4字节 - ID uint32 4字节  总共8字节
	//去看message.go 中定义的struct
	return 8
}

//封包
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte,error){

	dataBuf := bytes.NewBuffer([]byte{})

	//顺序：  len  id   data 不能搞混
	//将DataLen写进去dataBuf
	if err :=binary.Write(dataBuf,binary.LittleEndian,msg.GetDataLen());err != nil{
		return nil, err
	}

	//将Id写进去dataBuf
	if err :=binary.Write(dataBuf,binary.LittleEndian,msg.GetMsgId());err != nil{
		return nil, err
	}

	//将Data写进去dataBuf
	if err :=binary.Write(dataBuf,binary.LittleEndian,msg.GetData());err != nil{
		return nil, err
	}

	fmt.Println("封包源数据：",msg)
	fmt.Println("封包源数据：msg.Data",msg.GetData())
	fmt.Println("封包成功str：",string(dataBuf.Bytes()))
	fmt.Println("封包成功[]byte：",dataBuf.Bytes())

	return dataBuf.Bytes(),nil
}

//拆包
func (d *DataPack) Unpack(binaryData []byte) (ziface.IMessage,error){

	fmt.Println("拆包源数据：",binaryData)
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head信息，获取dateLen和Id
	msg := &Message{}

	//读len
	if err := binary.Read(dataBuff,binary.LittleEndian,&msg.DataLen);err != nil{
		return nil, err
	}
	//读id
	if err := binary.Read(dataBuff,binary.LittleEndian,&msg.Id);err != nil{
		return nil, err
	}

	//判断len是否已经超出了我们允许的包最大长度
	if (utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize){
		return nil,errors.New("too large msg data recv!")
	}

	fmt.Println("拆包成功：",msg)
	fmt.Println("拆包成功dataLen：",msg.DataLen)
	fmt.Println("拆包成功id：",msg.Id)
	fmt.Println("拆包成功[]byte：",dataBuff)
	return msg,nil
}
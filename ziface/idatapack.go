/**
 *@Desc:  粘包拆包
		解决TCP粘包问题，针对Message进行TLV格式的拆包
		长度 类型 内容
		先读取固定长度的head，消息内容的长度和消息的类型
		再根据消息内容的长度，再次进行一次读写，从conn中读取消息内容
 *@Author:Giousa
 *@Date:2021/1/14
 */
package ziface

type IDataPack interface {

	//获取包的头的长度方法
	GetHeadLen() uint32
	//封包
	Pack(msg IMessage) ([]byte,error)
	//拆包
	Unpack(binaryData []byte) (IMessage,error)
}

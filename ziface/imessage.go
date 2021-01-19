/**
 *@Desc: 封装消息
 *@Author:Giousa
 *@Date:2021/1/14
 */
package ziface

type IMessage interface {

	//获取消息ID
	GetMsgId() uint32
	//获取消息长度
	GetDataLen() uint32
	//获取消息内容
	GetData() []byte
	//设置消息ID
	SetMsgId(id uint32)
	//设置消息长度
	SetDataLen(len uint32)
	//设置消息内容
	SetData(data []byte)

}
/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/14
 */
package znet

type Message struct {
	Id      uint32 //消息id
	DataLen uint32 //消息长度
	Data    []byte //消息内容
}


//获取消息ID
func (m *Message) GetMsgId() uint32{

	return m.Id
}
//获取消息长度
func (m *Message) GetDataLen() uint32{

	return m.DataLen
}
//获取消息内容
func (m *Message) GetData() []byte{
	return m.Data
}
//设置消息ID
func (m *Message) SetMsgId(id uint32){
	m.Id = id
}
//设置消息长度
func (m *Message) SetDataLen(len uint32){
	m.DataLen = len
}
//设置消息内容
func (m *Message) SetData(data []byte){
	m.Data = data
}
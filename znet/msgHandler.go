/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/2/7
 */
package znet

import (
	"fmt"
	"strconv"
	"zinx/ziface"
)

type MsgHandler struct {

	//存放每个MsgId对应的处理方法
	Apis map[uint32]ziface.IRouter
}

//初始化或创建MsgHandler方法
func NewMsgHandler() *MsgHandler{
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}


//调度或执行对应Router消息处理方法
func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest){

	//从request中找到msgId
	hander,ok := mh.Apis[request.GetMsgID()]
	if !ok{
		fmt.Println("api msgId",request.GetMsgID(),"is not found!")
		return
	}

	//根据msgId，调度对应router业务
	hander.PreHandle(request)
	hander.Handle(request)
	hander.PostHandle(request)

}
//为消息添加具体处理器
func (mh *MsgHandler) AddRouter(msgID uint32,router ziface.IRouter){
	//判断当前msg绑定的处理方法是否已存在
	if _,ok := mh.Apis[msgID]; ok{
		//已存在
		fmt.Println("路由绑定失败！")
		panic("repeat api,msgId = "+strconv.Itoa(int(msgID)))
	}

	//添加
	mh.Apis[msgID] = router
	fmt.Println("路由绑定成功！")
}
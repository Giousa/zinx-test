/**
 *@Desc: 消息管理抽象层
 *@Author:Giousa
 *@Date:2021/2/7
 */
package ziface

type IMsgHandler interface {

	//调度或执行对应Router消息处理方法
	DoMsgHandler(request IRequest)
	//为消息添加具体处理器
	AddRouter(msgID uint32,router IRouter)
}

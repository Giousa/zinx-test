/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/5
 */
package ziface

type IRouter interface {
	//处理conn业务前的钩子方法Hook
	PreHandle(request IRequest)
	//处理conn业务主方法Hook
	Handle(request IRequest)
	//处理conn业务后钩子方法Hook
	PostHandle(request IRequest)
}

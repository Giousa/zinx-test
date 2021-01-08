/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/5
 */
package znet

import "zinx/ziface"

//实现router时，先嵌入BaseRouter基类，根据需要。对这个基类进行重写
type BaseRouter struct {

}


//处理conn业务前的钩子方法Hook
func (br *BaseRouter) PreHandle(request ziface.IRequest){

}
//处理conn业务主方法Hook
func (br *BaseRouter) Handle(request ziface.IRequest){

}
//处理conn业务后钩子方法Hook
func (br *BaseRouter) PostHandle(request ziface.IRequest){

}
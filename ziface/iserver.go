/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/3
 */
package ziface

type IServer interface {

	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Run()

	//路由
	AddRouter(router IRouter)

}
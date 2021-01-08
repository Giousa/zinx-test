/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/5
 */
package ziface

type IRequest interface {
	//得到当前链接
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte
}

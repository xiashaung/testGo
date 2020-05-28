package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var(
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin:func(r *http.Request) bool{
			return true
		},
	}
)

/**
  *  websocket处理程序
 */
func Wshandle(c *gin.Context)  {
	var(
		wsCoon *websocket.Conn //
		err error //错误消息
		data []byte //数据
		conn *Connection //
	)
	if wsCoon,err = upgrader.Upgrade(c.Writer,c.Request,nil);err !=nil {
		return
	}
	if conn,err = InitConnection(wsCoon);err!=nil {
		conn.Close();
	}
	//循环监听是否输入数据
	for  {
		if data,err = conn.ReadMessage();err!=nil {
			conn.Close()
		}
	}
}
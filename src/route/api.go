package route

import (
	"github.com/gin-gonic/gin"
	ws "github.com/my/repo/pkg/websocket"
)
import "github.com/my/repo/Controllers"

func InitRouter()*gin.Engine  {
	r := gin.New()
	gin.SetMode("debug")
	r.GET("/add",Controllers.Add)
	r.GET("/ws",ws.Wshandle)
	r.GET("/del",Controllers.Del)
	r.POST("/createCard",Controllers.CreateCard)
	r.GET("/addUser",Controllers.AddUser)
	return r
}


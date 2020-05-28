package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/pkg/http"
)

var Response *response

type response struct {

}

func init()  {
	Response = new(response)
}

func (res response)ErrorJson(c *gin.Context,errmsg string){
	c.JSON(200,gin.H{
		"code":http.API_ERROR,
		"msg":errmsg,
	})
}

func (res response)SuccessJson(c *gin.Context,data interface{}){
	c.JSON(200,gin.H{
		"code":http.API_SUCCESS,
		"msg":"ok",
		"data":data,
	})
}

func (res response)EmptySuccessJson(c *gin.Context){
	c.JSON(200,gin.H{
		"code":http.API_SUCCESS,
		"msg":"ok",
		"data":"[]",
	})
}

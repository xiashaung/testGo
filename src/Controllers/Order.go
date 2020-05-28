package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/my/repo/Models"
	"github.com/my/repo/pkg"
)

type order struct {
}

func Add(c *gin.Context) {
	var card []Models.Card
	pkg.DB.Model(Models.Card{}).Select("*").Limit(10).Scan(&card)
	c.JSON(200, gin.H{
		"message": "message",
		"data":    card,
	})
}

func Del(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "dataname",
	})
}

func CreateCard(c *gin.Context){
	card_no,_:= c.GetPostForm("card_no")
	if card_no == "" {
		pkg.Response.ErrorJson(c,"编号不能为空")
	}else{
		card := Models.Card{Card_no: card_no}
		pkg.DB.Model(Models.Card{}).Create(&card)
		pkg.Response.SuccessJson(c,&card)
	}

}

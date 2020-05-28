package Models

import (
	"github.com/jinzhu/gorm"
	"github.com/my/repo/pkg"
)

type Card struct {
	Id      uint   `json:"id" form:"id"`
	Card_no string `json:"card_no" form:"card_no"`
}

func (Card)TableName()string {
  return "xs_card"
}

func (Card) instance()*gorm.DB{
   return pkg.DB.Model(Card{});
}
package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    int64   `gorm:"column:user_id" json:"user_id"`
	ProductId int64   `gorm:"column:product_id" json:"product_id"`
	Num       int64   `gorm:"column:num" json:"num"`
	Status    int     `gorm:"column:status" json:"status"`
	Amount    float64 `gorm:"column:amount" json:"amount"`
}

func (o *Order) TableName() string {
	return "ts_order"
}

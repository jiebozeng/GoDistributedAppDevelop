package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string  `gorm:"column:product_name" json:"product_name"`
	InvenNum    int64   `gorm:"column:inven_num" json:"inven_num"`
	Status      int     `gorm:"column:status" json:"status"`
	Price       float64 `gorm:"column:price" json:"price"`
}

// 表名
func (p *Product) TableName() string {
	return "ts_product"
}

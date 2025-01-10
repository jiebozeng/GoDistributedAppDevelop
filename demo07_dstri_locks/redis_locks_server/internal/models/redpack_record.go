package models

import "gorm.io/gorm"

type RedPackRecord struct {
	gorm.Model
	RedPackId int64  `gorm:"column:redpack_id" json:"redpack_id"`
	UserId    int64  `gorm:"column:user_id" json:"user_id"`
	UserName  string `gorm:"column:user_name" json:"user_name"`
	Amount    int64  `gorm:"column:amount" json:"amount"`
}

func (r *RedPackRecord) TableName() string {
	return "ts_redpack_record"
}

package models

import "gorm.io/gorm"

type Redpack struct {
	gorm.Model
	Amount    int64 `gorm:"column:amount" json:"amount"`
	Num       int64 `gorm:"column:num" json:"num"`
	ValidTime int64 `gorm:"column:valid_time" json:"valid_time"`
	Status    int   `json:"status"    orm:"status"     dc:"状态:1可用,2已结束,3已取消,5已领完"`
	ProNum    int64 `json:"pro_num"   orm:"pro_num"   dc:"已经领取数量"`
}

func (r *Redpack) TableName() string {
	return "ts_redpack"
}

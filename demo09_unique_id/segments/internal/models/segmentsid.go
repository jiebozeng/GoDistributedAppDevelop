package models

import "gorm.io/gorm"

type SegmentsId struct {
	gorm.Model
	Id        int64  `gorm:"column:id" json:"id"`
	MaxId     int64  `gorm:"column:max_id" json:"max_id"`     //当前最大id
	Step      int64  `gorm:"column:step" json:"step"`         //号段的步长
	BizType   int64  `gorm:"column:biz_type" json:"biz_type"` //业务类型
	Version   int64  `gorm:"column:version" json:"version"`   // 版本号
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

func (s *SegmentsId) TableName() string {
	return "id_generator"
}

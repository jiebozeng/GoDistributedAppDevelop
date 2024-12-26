package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId     int64  `gorm:"column:user_id" json:"user_id"`
	UserName   string `gorm:"column:user_name" json:"user_name"`     //用户名
	UserPwd    string `gorm:"column:user_pwd" json:"user_pwd"`       //密码
	UserMobile string `gorm:"column:user_mobile" json:"user_mobile"` //手机号码
	UserEmail  string `gorm:"column:user_email" json:"user_email"`   //邮箱
	CreatedAt  string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  string `gorm:"column:updated_at" json:"updated_at"`
}

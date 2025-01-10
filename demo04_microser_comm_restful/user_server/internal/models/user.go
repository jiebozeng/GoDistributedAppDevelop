package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId     int64  `gorm:"column:user_id" json:"user_id"`
	UserName   string `gorm:"column:user_name" json:"user_name"`     //用户名
	UserPwd    string `gorm:"column:user_pwd" json:"user_pwd"`       //密码
	UserMobile string `gorm:"column:user_mobile" json:"user_mobile"` //手机号码
}

func (u *User) TableName() string {
	return "ts_user"
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId    int64  `json:"userid" gorm:"type:bigint;column:userid"`
	UserName  string `json:"username" gorm:"type:varchar(200);column:username"  binding:"required,min=3,max=20"`
	PassWord  string `json:"password" gorm:"type:varchar(200);column:password"  binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
	Email     string `gorm:"column:email"`
}

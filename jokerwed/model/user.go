package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId   int64  `json:"user_id" gorm:"type:bigint;column:userid"`
	UserName string `json:"username" gorm:"type:varchar(200);column:username"`
	PassWord string `json:"password" gorm:"type:varchar(20);column:password"`
	Email    string `gorm:"column:email"`
}

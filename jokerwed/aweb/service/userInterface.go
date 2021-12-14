package service

import (
	model2 "jokerweb/model"
)

type UserInterface interface {
	Login(user *model2.User) (token string, err error)
	Register(user *model2.User) (err error)
}

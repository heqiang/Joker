package dao

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"jokerweb/global"
	"jokerweb/model"
	"jokerweb/utils/jwt"
	"strings"
)

func QueryByUser(userName, pwd string) (token string, err error) {
	var user model.User
	queryRes := global.Db.Where("username=?", userName).Take(&user)
	if queryRes.RowsAffected == 0 {
		err = errors.New("用户不存在")
		return
	}
	passwordInfo := strings.Split(user.PassWord, "$")
	options := &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	check := password.Verify(pwd, passwordInfo[2], passwordInfo[3], options)
	if !check {
		err = errors.New("密码错误")
		return
	}
	return jwt.GenToken(user.UserName, user.UserId)
}

func QueryUserByName(username string) error {
	var user model.User
	res := global.Db.Where("username=?", username).Take(&user)
	if res.RowsAffected != 0 {
		return errors.New("已存在该用户")
	}
	return nil
}

func InsertUser(userinfo *model.User) error {
	options := &password.Options{SaltLen: 10, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(userinfo.PassWord, options)
	userinfo.PassWord = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	res := global.Db.Create(userinfo)
	if res.RowsAffected >= 1 {
		return nil
	}
	return res.Error

}

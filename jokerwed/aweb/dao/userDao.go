package dao

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"jokerweb/global"
	"jokerweb/model"
	"jokerweb/utils/jwt"
)

const serct = "1422127065@qq.com"

// 用户密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(serct))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func QueryByUser(userName, password string) (token string, err error) {
	var user model.User
	queryRes := global.Db.Where("username=?", userName).Take(&user)
	if queryRes.RowsAffected == 0 {
		err = errors.New("用户不存在")
		return
	}
	if user.PassWord != encryptPassword(password) {
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
	userinfo.PassWord = encryptPassword(userinfo.PassWord)
	res := global.Db.Create(userinfo)
	if res.RowsAffected >= 1 {
		return nil
	}
	return res.Error

}

func QueryByUserId(id int64) (user *model.User) {
	user = new(model.User)
	global.Db.Where("userid=?", id).Take(&user)
	return
}

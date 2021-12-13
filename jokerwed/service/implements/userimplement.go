package implements

import (
	"jokerweb/dao"
	"jokerweb/model"
	"jokerweb/utils/snowflake"
)

type User struct {
}

func (u *User) Login(user *model.User) (token string, err error) {
	return dao.QueryByUser(user.UserName, user.PassWord)

}
func (u *User) Register(user *model.User) (err error) {
	err = dao.QueryUserByName(user.UserName)
	if err != nil {
		return
	}
	// 生成Uid
	UUID := snowflake.GetSnowId()
	userinfo := model.User{
		UserId:   UUID,
		UserName: user.UserName,
		PassWord: user.PassWord,
	}
	// mysql入库
	return dao.InsertUser(&userinfo)
}

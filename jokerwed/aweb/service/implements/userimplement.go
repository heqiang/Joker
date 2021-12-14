package implements

import (
	dao2 "jokerweb/aweb/dao"
	model2 "jokerweb/model"
	"jokerweb/utils/snowflake"
)

type User struct {
}

func (u *User) Login(user *model2.User) (token string, err error) {
	return dao2.QueryByUser(user.UserName, user.PassWord)

}
func (u *User) Register(user *model2.User) (err error) {
	err = dao2.QueryUserByName(user.UserName)
	if err != nil {
		return
	}
	// 生成Uid
	UUID := snowflake.GetSnowId()
	userinfo := model2.User{
		UserId:   UUID,
		UserName: user.UserName,
		PassWord: user.PassWord,
	}
	// mysql入库
	return dao2.InsertUser(&userinfo)
}

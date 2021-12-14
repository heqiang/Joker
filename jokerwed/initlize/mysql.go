package initlize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	model2 "jokerweb/aweb/model"
	"jokerweb/config"
	"jokerweb/global"
)

func InitMysql(conf *config.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("数据库打开失败,err:%s", err)
		panic(err)
	}
	var article model2.Article
	var user model2.User
	var vote model2.Vote
	err = db.AutoMigrate(&article, &user, &vote)
	if err != nil {
		fmt.Println("数据库迁移失败")
		return
	}
	global.Db = db
	return nil
}

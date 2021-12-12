package initlize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"jokerweb/config"
	"jokerweb/global"
	"jokerweb/model"
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
	var article model.Article
	var user model.User
	var vote model.Vote
	err = db.AutoMigrate(&article, &user, &vote)
	if err != nil {
		fmt.Println("数据库迁移失败")
		return
	}
	global.Db = db
	return nil
}

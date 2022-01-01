package initlize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		fmt.Println(dsn)
		fmt.Printf("数据库打开失败,err:%s", err)
		panic(err)
	}
	fmt.Println("mysql:", dsn)
	var article model.Article
	var user model.User
	var comment model.ArticleComment
	err = db.AutoMigrate(&article, &user, &comment)
	if err != nil {
		fmt.Println("数据库迁移失败")
		return
	}
	global.Db = db
	return nil
}

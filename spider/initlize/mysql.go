package initlize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"spider/global"
	"spider/model"
)

func InitMysql() {
	dsn := "root:142212@tcp(127.0.0.1:3306)/jokeweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("数据库打开失败,err:%s", err)
		panic(err)
	}
	var spideritem model.SpiderItem
	err = db.AutoMigrate(&spideritem)
	if err != nil {
		fmt.Println("数据库迁移失败")
		return
	}
	global.MySqlDb = db
}

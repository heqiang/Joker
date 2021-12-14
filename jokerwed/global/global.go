package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"jokerweb/config"
)

var (
	Db   *gorm.DB
	Conf *config.AppConfig
	Rdb  *redis.Client
	Es   *elastic.Client
)

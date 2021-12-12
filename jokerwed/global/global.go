package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"jokerweb/config"
)

var (
	Db   *gorm.DB
	Conf *config.AppConfig
	Rdb  *redis.Client
)

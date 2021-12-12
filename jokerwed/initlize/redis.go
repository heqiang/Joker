package initlize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"jokerweb/config"
	"time"
)

func InitRedis(conf *config.RedisConfig) (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password, // no password set
		DB:       conf.Db,       // use default DB
		PoolSize: conf.PoolSize, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil

}

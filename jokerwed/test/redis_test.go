package test

import (
	"context"
	"fmt"
	"jokerweb/global"
	"jokerweb/initlize"
	"testing"
)

func Test_redis(t *testing.T) {
	var ctx context.Context
	err := initlize.InitConfig()
	if err != nil {
		fmt.Println("初始化配置出错")
		return
	}
	err = initlize.InitRedis(global.Conf.RedisConfig)
	if err != nil {
		fmt.Println("redis初始化错误")
		return
	}
	res := global.Rdb.ZScore(ctx, "dddd", "ssss")
	fmt.Println(res.Err())
	fmt.Println(res.Result())
}

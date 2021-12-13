package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"jokerweb/controller"
	"jokerweb/global"
	"jokerweb/initlize"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := initlize.InitConfig()
	if err != nil {
		fmt.Println("初始化配置出错")
		return
	}
	err = initlize.InitLogger(global.Conf.LogConfig, "warn")
	if err != nil {
		fmt.Println("日志配置初始化错误")
		return
	}
	defer zap.L().Sync()
	err = initlize.InitMysql(global.Conf.MysqlConfig)
	if err != nil {
		fmt.Println("mysql初始化错误")
		return
	}
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Println("翻译器初始化失败")
		return
	}
	//err = initlize.InitRedis(global.Conf.RedisConfig)
	//if err != nil {
	//	fmt.Println("redis初始化错误")
	//	return
	//}

	//开启服务
	r := initlize.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("app.port")),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	log.Println("Server exiting")
}

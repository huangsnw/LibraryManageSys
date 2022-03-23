package main

import (
	"LibraryManageSys/pkg/logs"
	"LibraryManageSys/pkg/result"
	"LibraryManageSys/pkg/setting"
	"LibraryManageSys/pkg/util"
	"LibraryManageSys/routers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 读取配置文件
	setting.ConfigViper()
	// 读取日志配置
	logs.InitLog()
	// 初始化数据库
	util.InitDB()
	// 初始化Redis
	util.InitRedis()
	// 初始化RabbitMQ
	// util.InitRabbitMQ()
	// 定时任务
	// task.PrintTime()
	// 开启gin
	engine := gin.Default()
	// 消息推送
	// util.SendMessage(util.RabbitChannel, "Hello World1")
	// 消息接收
	// util.GetMessage(util.RabbitChannel)
	// RabbitMQ关闭
	// util.CloseRabbitMQ()
	// 拦截器
	engine.Use(TokenHandle)
	// 路由
	engine = routers.CollectRouter(engine)
	fmt.Println("---------------------------")
	fmt.Println("http://" + viper.GetString("server.host") + ":" + viper.GetString("server.port") + "/swagger/index.html")
	fmt.Println("---------------------------")
	// 运行gin
	err := engine.Run(":" + viper.GetString("server.port"))
	if err != nil {
		log.Fatalf("Gin 框架出错")
	}
}

func TokenHandle(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("X-Token")
	_, err := util.ParseToken(tokenString)
	res := result.Result{}
	if err != nil && ctx.FullPath() != "/token/get" && ctx.FullPath() != "/swagger/*any" {
		res.Error("鉴权失败")
		ctx.JSON(500, res)
		ctx.Abort()
		log.Print("Token 鉴权失败")
		return
	} else {
		ctx.Next()
	}
}

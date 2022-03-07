/*
 * @title: 这里写标题
 * @Date: 2022-02-15 17:17:38
 * @version: 1.0
 * @author: huang sn
 * @description: 这里写描述信息
 * @FilePath: /LibraryManageSys/main.go
 */
package main

import (
	"LibraryManageSys/pkg/logs"
	"LibraryManageSys/pkg/setting"
	"LibraryManageSys/pkg/util"
	"LibraryManageSys/routers"
	"fmt"
	"log"
	"time"

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
	// 开启gin
	engine := gin.Default()
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

/**
 * @description:
 * @Accept:
 * @param {*gin.Context} ctx
 * @return {*}
 * @Router:
 */
func TokenHandle(ctx *gin.Context) {
	tokenString := ctx.Request.Header.Get("X-Token")
	_, err := util.ParseToken(tokenString)
	if err != nil && ctx.FullPath() != "/token/get" && ctx.FullPath() != "/swagger/*any" {
		ctx.JSON(500, gin.H{
			"code":      500,
			"message":   "鉴权失败",
			"timestamp": time.Now(),
		})
		ctx.Abort()
		log.Print("Token 鉴权失败")
		return
	} else {
		ctx.Next()
	}
}

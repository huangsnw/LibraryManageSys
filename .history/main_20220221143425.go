/*
 * @BasePath: main.go
 * @Summary: Do not edit
 * @Schemes: Do not edit
 * @Description:
 * @Tags: your name
 * @Accept: Do not edit
 * @Produce:
 * @Success:
 * @Router:
 */
package main

import (
	"LibraryManageSys/app/routers"
	"LibraryManageSys/storage/logs"
	"LibraryManageSys/util"
	"log"

	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	// swagger embed files

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 读取配置文件
	util.ConfigViper()
	// 读取日志配置
	logs.InitLog()
	// 初始化数据库
	util.InitDB()
	// 初始化Redis
	util.InitDB()
	// 开启gin
	engine := gin.Default()
	// 路由
	engine = routers.CollectRouter(engine)
	// 运行gin
	err := engine.Run(":" + viper.GetString("server.port"))
	if err != nil {
		log.Fatalf("Gin 框架出错")
	}
}

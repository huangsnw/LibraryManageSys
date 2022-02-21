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

	"github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title 这里写标题
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath 这里写base path
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

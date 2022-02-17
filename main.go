package main

import (
	"LibraryManageSys/app/routers"
	"LibraryManageSys/util"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 读取配置文件
	util.ConfigViper()
	// 初始化数据库
	db := util.InitDB
	fmt.Println(db)
	// 开启gin
	engine := gin.Default()
	// 路由
	engine = routers.CollectRouter(engine)
	// 运行gin
	err := engine.Run()
	if err != nil {
		log.Fatalf("Gin 框架出错")
	}
}

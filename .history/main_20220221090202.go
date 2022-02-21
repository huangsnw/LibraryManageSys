package main

import (
	"os"
)

func main() {
	// // 读取配置文件
	// util.ConfigViper()
	// // 读取日志配置
	// logs.InitLog()
	// // 初始化数据库
	// util.InitDB()
	// // 开启gin
	// engine := gin.Default()
	// // 路由
	// engine = routers.CollectRouter(engine)
	// // 运行gin
	// err := engine.Run(":" + viper.GetString("server.port"))
	// if err != nil {
	// 	log.Fatalf("Gin 框架出错")
	// }

	path := "test.txt"
	file := os.Open(path)
}

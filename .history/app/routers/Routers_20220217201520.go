package routers

import (
	"LibraryManageSys/app/controllers"

	"github.com/gin-gonic/gin"
)

func CollectRouter(engine *gin.Engine) *gin.Engine {
	// 中间件
	engine.Use()

	// 用户
	userGroup := engine.Group("/user")
	// 新增用户
	userGroup.POST("/save",controllers.)
	// 查询用户
	userGroup.GET("/select")
	// 删除用户
	userGroup.DELETE("/del")
	// 修改用户
	userGroup.PUT("/update")

	// 图书
	bookGroup := engine.Group("/book")
	// 新增图书
	bookGroup.POST("/save", controllers.SaveBook)
	// 查询图书
	bookGroup.GET("/select", controllers.SelectBook)
	// 删除图书
	bookGroup.DELETE("/del", controllers.DeleteBook)
	// 修改图书
	bookGroup.POST("/update", controllers.UpdateBook)

	// 返回值
	return engine
}

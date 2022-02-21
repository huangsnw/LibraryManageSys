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
	userGroup.POST("/save", controllers.SaveUser)
	// 查询用户
	userGroup.GET("/select", controllers.SelectUser)
	// 删除用户
	userGroup.DELETE("/del", controllers.DeleteUser)
	// 修改用户
	userGroup.PUT("/update", controllers.UpdateUser)
	// 导出用户
	userGroup.GET("/export")

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

	// 文件
	fileUpload := engine.Group("/file")
	// 单个文件上传
	fileUpload.POST("/upload", controllers.UploadFile)
	// 多个文件上传
	fileUpload.POST("/uploads", controllers.UploadFiles)

	// 返回值
	return engine
}

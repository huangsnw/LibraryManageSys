package routers

import (
	_ "LibraryManageSys/docs"
	"LibraryManageSys/middleware"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func CollectRouter(engine *gin.Engine) *gin.Engine {
	// 中间件
	engine.Use()

	// swagger 路由
	engine.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 用户
	userGroup := engine.Group("/user")
	// 新增用户
	userGroup.POST("/save", middleware.SaveUser)
	// 查询用户
	userGroup.GET("/select", middleware.SelectUser)
	// 删除用户
	userGroup.DELETE("/del", middleware.DeleteUser)
	// 修改用户
	userGroup.PUT("/update", middleware.UpdateUser)
	// 导出用户
	userGroup.GET("/export", middleware.ExportUser)
	// 用户分页查询
	userGroup.GET("/page", middleware.PageSelectUser)

	// 图书
	bookGroup := engine.Group("/book")
	// 新增图书
	bookGroup.POST("/save", middleware.SaveBook)
	// 查询图书
	bookGroup.GET("/select", middleware.SelectBook)
	// 删除图书
	bookGroup.DELETE("/del", middleware.DeleteBook)
	// 修改图书
	bookGroup.POST("/update", middleware.UpdateBook)
	// 图书分页查询
	bookGroup.GET("/page", middleware.PageSelectBook)
	// 查询全部图书
	bookGroup.GET("/all", middleware.SelectAllBook)

	// 文件
	fileUpload := engine.Group("/file")
	// 单个文件上传
	fileUpload.POST("/upload", middleware.UploadFile)
	// 多个文件上传
	fileUpload.POST("/uploads", middleware.UploadFiles)

	// token 验证
	token := engine.Group("/token")
	// 拿到token
	token.GET("/get", middleware.GetToken)

	// 测试
	ginTest := engine.Group("/test")
	ginTest.GET("/one", middleware.TestASCII)

	// 随机图片API
	randomImages := engine.Group("/getImage")
	randomImages.GET("/api", middleware.Get)

	// 返回值
	return engine
}

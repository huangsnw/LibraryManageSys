/*
 * @title: 这里写标题
 * @Date: 2022-02-16 17:49:46
 * @version: 1.0
 * @author: huang sn
 * @description: 这里写描述信息
 * @FilePath: /LibraryManageSys/app/routers/Routers.go
 */
/*
 * @title: 这里写标题
 * @version: 1.0
 * @description: 这里写描述信息
 * @BasePath: 这里写base path
 */
package routers

import (
	"LibraryManageSys/app/controllers"

	_ "LibraryManageSys/docs"

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
	userGroup.POST("/save", controllers.SaveUser)
	// 查询用户
	userGroup.GET("/select", controllers.SelectUser)
	// 删除用户
	userGroup.DELETE("/del", controllers.DeleteUser)
	// 修改用户
	userGroup.PUT("/update", controllers.UpdateUser)
	// 导出用户
	userGroup.GET("/export", controllers.ExportUser)

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

	// token 验证
	token := engine.Group("/token")
	// 拿到token
	token.GET("/get", controllers.GetToken)

	// 返回值
	return engine
}

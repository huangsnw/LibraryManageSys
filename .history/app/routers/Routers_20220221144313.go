/*
 * @Author: your name
 * @Date: 2022-02-16 17:49:46
 * @LastEditTime: 2022-02-21 14:43:13
 * @LastEditors: your name
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /LibraryManageSys/app/routers/Routers.go
 */
package routers

import (
	"LibraryManageSys/app/controllers"

	_ "bluebell/docs" // 千万不要忘了导入把你上一步生成的docs

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

	// 返回值
	return engine
}

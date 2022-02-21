/*
 * @Author: your name
 * @Date: 2022-02-16 16:21:41
 * @LastEditTime: 2022-02-21 14:42:06
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /LibraryManageSys/app/controllers/BookController.go
 */
package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200
// @Router /posts2 [get]
func SaveBook(ctx *gin.Context) {
	var book dao.Book
	ctx.BindJSON(&book)
	fmt.Println(util.DB.Name())
	util.DB = util.DB.Create(&book)
	affected := util.DB.RowsAffected
	fmt.Println("受影响的行：", affected)
	if affected == 1 {
		ctx.JSON(200, gin.H{
			"message":   "sucess",
			"timestamp": time.Now(),
		})
	} else {
		ctx.JSON(500, gin.H{
			"message": "error",
			"date":    time.Now(),
		})
	}
}

func SelectBook(ctx *gin.Context) {
	var book dao.Book
	bookId := ctx.Query("id")
	util.DB = util.DB.Take(&book, "id = ?", bookId)
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, book)
	} else {
		ctx.JSON(500, gin.H{
			"message":   "error",
			"timestamp": time.Now(),
		})
	}
}

func UpdateBook(ctx *gin.Context) {
	var book dao.Book
	ctx.BindJSON(&book)
	util.DB = util.DB.Model(&book).Updates(&book)
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, gin.H{
			"message":   "sucess",
			"timestamp": time.Now(),
			"data":      book,
		})
	} else {
		ctx.JSON(500, gin.H{
			"message":   "error",
			"timestamp": time.Now(),
			"data":      nil,
		})
	}
}

func DeleteBook(ctx *gin.Context) {
	var book dao.Book
	bookId := ctx.Query("id")
	util.DB = util.DB.Delete(&book, "id = ?", bookId)
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, gin.H{
			"message":   "sucess",
			"timestamp": time.Now(),
			"data":      nil,
		})
	}
}

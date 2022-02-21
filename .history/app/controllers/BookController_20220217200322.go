package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// 保存图书
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

// 查询图书
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

// 更新图书
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

// 删除图书
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

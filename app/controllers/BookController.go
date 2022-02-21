/*
 * @title: 这里写标题
 * @version: 1.0
 * @author: huang sn
 * @description: 这里写描述信息
 * @BasePath: 这里写base path
 */
package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @description:
 * @param {*gin.Context} ctx
 * @return {*}
 */
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

/**
 * @description:
 * @param {*gin.Context} ctx
 * @return {*}
 */
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

/**
 * @description:
 * @param {*gin.Context} ctx
 * @return {*}
 */
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

/**
 * @description:
 * @param {*gin.Context} ctx
 * @return {*}
 */
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

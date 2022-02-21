package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

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

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
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

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
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

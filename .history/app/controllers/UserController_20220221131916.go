package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"
	"time"

	"github.com/gin-gonic/gin"
)

// GetPostListHandler2 升级版帖子列表接口
// @Summary 保存用户
// @Description 保存用户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /save [get]
func SaveUser(ctx *gin.Context) {
	var user dao.User
	ctx.BindJSON(&user)
	util.DB.AutoMigrate(&dao.User{})
	util.DB = util.DB.Create(&user)
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, gin.H{
			"message":   "sucess",
			"timestamp": time.Now(),
			"data":      user,
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
func DeleteUser(ctx *gin.Context) {
	var user dao.User
	userId := ctx.Query("id")
	util.DB = util.DB.Delete(&user, "id = ? ", userId)
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, gin.H{
			"message":   "sucess",
			"timestamp": time.Now(),
			"data":      nil,
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
func SelectUser(ctx *gin.Context) {
	var user dao.User
	userId := ctx.Query("id")
	util.DB = util.DB.Take(&user, "id = ?", userId)
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, gin.H{
			"message":   "success",
			"timestamp": time.Now(),
			"data":      user,
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
func UpdateUser(ctx *gin.Context) {
	var user dao.User
	ctx.BindJSON(&user)
	util.DB = util.DB.Model(&user).Updates(&user)
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, gin.H{
			"message":   "success",
			"timestamp": time.Now(),
			"data":      user,
		})
	} else {
		ctx.JSON(500, gin.H{
			"message":   "error",
			"timestamp": time.Now(),
			"data":      nil,
		})
	}
}

func ExportUser(ctx *gin.Context) {
	var user dao.User
	var arrayUser []dao.User
	result := util.DB.Find(&user)
	arrayUser := util.DB.RowsAffected
}

package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"
	"time"

	"github.com/gin-gonic/gin"
)

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
	util.DB = util.DB.Find(&user)
}

func SelectUserByPage(ctx *gin.Context)

package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"
	"time"

	"github.com/gin-gonic/gin"
)

// 保存用户
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

// 删除用户

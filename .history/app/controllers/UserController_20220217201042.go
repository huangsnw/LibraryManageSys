package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"
	"time"

	"github.com/gin-gonic/gin"
)

// 保存用户
func UserSave(ctx *gin.Context) {
	var user dao.User
	ctx.BindJSON(&user)
	util.DB = util.DB.Create(&user)
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, gin.H{
			"message":   "sucess",
			"timestamp": time.Now(),
			"data":      user,
		})
	} else {

	}

}

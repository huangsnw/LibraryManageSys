package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"

	"github.com/gin-gonic/gin"
)

// 保存用户
func UserSave(ctx *gin.Context) {
	var user dao.User
	ctx.BindJSON(&user)
	util.DB = util.DB.Create(&user)

}

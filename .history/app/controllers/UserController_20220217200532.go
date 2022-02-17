package controllers

import (
	"LibraryManageSys/dao"

	"github.com/gin-gonic/gin"
)

// 保存用户
func UserSave(ctx *gin.Context) {
	var user dao.User
	ctx.Query(&user)

}

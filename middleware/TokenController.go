package middleware

import (
	"LibraryManageSys/pkg/result"
	"LibraryManageSys/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 获取Token
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /token/get [get]
func GetToken(ctx *gin.Context) {
	name := ctx.Query("username")
	password := ctx.Query("password")
	res := result.Result{}
	if name == "admin" && password == "123" {
		ctx.JSON(http.StatusOK, res.SUCESS(util.GetToken(name)))
	} else {
		ctx.JSON(http.StatusInternalServerError, res.Error("登录失败"))
	}
}

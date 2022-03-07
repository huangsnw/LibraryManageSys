package middleware

import (
	"LibraryManageSys/models"
	"LibraryManageSys/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func SaveUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	util.DB.AutoMigrate(&models.User{})
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

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func DeleteUser(ctx *gin.Context) {
	var user models.User
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

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func SelectUser(ctx *gin.Context) {
	var user models.User
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

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func UpdateUser(ctx *gin.Context) {
	var user models.User
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

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func ExportUser(ctx *gin.Context) {
	var user models.User
	util.DB = util.DB.Find(&user)
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func SelectUserByPage(ctx *gin.Context) {

}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func GetToken(ctx *gin.Context) {
	name := ctx.Query("username")
	password := ctx.Query("password")
	if name == "admin" && password == "123" {
		result := models.Result{
			Code:    200,
			Message: util.GetToken(name),
			Data:    time.Now(),
		}
		ctx.JSON(http.StatusOK, result)
	} else {
		result := models.Result{
			Code:    500,
			Message: "登录失败",
			Data:    time.Now(),
		}
		ctx.JSON(500, result)
	}
}

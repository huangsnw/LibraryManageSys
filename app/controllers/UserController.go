/*
 * @title: 这里写标题
 * @Date: 2022-02-17 20:03:35
 * @version: 1.0
 * @author: huang sn
 * @description: 这里写描述信息
 * @FilePath: /LibraryManageSys/app/controllers/UserController.go
 */
package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/models"
	"LibraryManageSys/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/**
 * @description:
 * @Accept:
 * @param {*gin.Context} ctx
 * @return {*}
 * @Router:
 */
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

/**
 * @description:
 * @Accept:
 * @param {*gin.Context} ctx
 * @return {*}
 * @Router:
 */
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

/**
 * @description:
 * @Accept:
 * @param {*gin.Context} ctx
 * @return {*}
 * @Router:
 */
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

/**
 * @description:
 * @Accept:
 * @param {*}
 * @return {*}
 * @Router:
 */
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

/**
 * @description:
 * @Accept:
 * @param {*gin.Context} ctx
 * @return {*}
 * @Router:
 */
func ExportUser(ctx *gin.Context) {
	var user dao.User
	util.DB = util.DB.Find(&user)
}

/**
 * @description:
 * @Accept:
 * @param {*gin.Context} ctx
 * @return {*}
 * @Router:
 */
func SelectUserByPage(ctx *gin.Context) {

}

/**
 * @description:
 * @Accept:
 * @param {*gin.Context} ctx
 * @return {*}
 * @Router:
 */
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

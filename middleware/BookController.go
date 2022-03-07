package middleware

import (
	"LibraryManageSys/models"
	"LibraryManageSys/pkg/util"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary 保存图书
// @Produce  json
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func SaveBook(ctx *gin.Context) {
	var book models.Book
	ctx.BindJSON(&book)
	util.DB = util.DB.Model(&book).Create(&book).Updates(map[string]interface{}{"create_at": time.Now(), "create_by": util.Get("username")})
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

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func SelectBook(ctx *gin.Context) {
	bk := ctx.MustGet("id")
	log.Println(bk)
	var book models.Book
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

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func UpdateBook(ctx *gin.Context) {
	var book models.Book
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

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func DeleteBook(ctx *gin.Context) {
	var book models.Book
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

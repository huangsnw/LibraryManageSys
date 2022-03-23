package middleware

import (
	"LibraryManageSys/models"
	"LibraryManageSys/pkg/page"
	"LibraryManageSys/pkg/result"
	"LibraryManageSys/pkg/util"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @Summary 保存图书
// @Produce  json
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /book/save [post]
func SaveBook(ctx *gin.Context) {
	var book models.Book
	ctx.ShouldBind(&book)
	util.DB = util.DB.Create(&book)
	affected := util.DB.RowsAffected
	fmt.Println("受影响的行：", affected)
	res := result.Result{}
	if affected == 1 {
		ctx.JSON(200, res.SUCESS(&book))
	} else {
		ctx.JSON(500, res.ERROR())
	}
}

// @Summary 根据ID选择图书
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /book/select [get]
func SelectBook(ctx *gin.Context) {
	bk := ctx.MustGet("id")
	log.Println(bk)
	var book models.Book
	bookId := ctx.Query("id")
	util.DB = util.DB.Take(&book, "id = ?", bookId)
	res := result.Result{}
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, res.SUCESS(&book))
	} else {
		ctx.JSON(500, res.ERROR())
	}
}

// @Summary 根据ID修改图书
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /book/update [post]
func UpdateBook(ctx *gin.Context) {
	var book models.Book
	ctx.BindJSON(&book)
	util.DB = util.DB.Model(&book).Where("deleted = ?", 0).Updates(models.Book{
		BookName:       book.BookName,
		Author:         book.Author,
		Publisher:      book.Publisher,
		Isbn:           book.Isbn,
		Classification: book.Classification,
		Floor:          book.Floor,
		Bookshelf:      book.Bookshelf,
		DateOfPurchase: book.DateOfPurchase,
		Remarks:        book.Remarks,
		Picture:        book.Picture,
	})
	res := result.Result{}
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, res.OK())
	} else {
		ctx.JSON(500, res.ERROR())
	}
}

// @Summary 根据ID删除图书
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /book/del [delete]
func DeleteBook(ctx *gin.Context) {
	var book models.Book
	bookId := ctx.Query("id")
	util.DB = util.DB.Delete(&book, "id = ?", bookId)
	res := result.Result{}
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, res.OK())
	} else {
		ctx.JSON(500, res.ERROR())
	}
}

// @Summary 图书分页查询
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /book/page [get]
func PageSelectBook(ctx *gin.Context) {
	var book []models.Book
	page1, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", viper.GetString("page.pageSize")))
	numb, _ := strconv.Atoi(ctx.DefaultQuery("pageNumber", viper.GetString("page.pageNumber")))
	var sum int64
	util.DB.Model(&book).Count(&sum)
	err := util.DB.Limit(page1).Offset((numb - 1) * page1).Find(&book).Order("id desc").Error
	page2 := page.Page{}
	if err != nil {
		res := result.Result{}
		ctx.JSON(500, res.ERROR())
		return
	}
	page2.OK(page1, numb, int(sum), &book)
	ctx.JSON(200, page2)
}

func SelectAllBook(ctx *gin.Context) {
	var book []models.Book
	util.DB.Raw("select * from books where id is not null").Scan(&book)
	res := result.Result{}
	ctx.JSON(http.StatusOK, res.SUCESS(&book))
}

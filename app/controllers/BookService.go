package controllers

import (
	"LibraryManageSys/dao"
	"LibraryManageSys/util"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Save(ctx *gin.Context) {
	var book dao.Book
	err := ctx.BindJSON(&book)
	fmt.Println(book)
	fmt.Println(util.DB.Name())
	util.DB = util.DB.Create(&book)
	fmt.Println(book.ID)
	fmt.Println(book.BookName)
	fmt.Println(book.Author)
	fmt.Println(book.Bookshelf)
	fmt.Println(book.CreateAt)
	fmt.Println(book.CreateBy)
	affected := util.DB.RowsAffected
	if err != nil || affected != 1 {
		log.Fatalf("图书保存出错")
	}
}

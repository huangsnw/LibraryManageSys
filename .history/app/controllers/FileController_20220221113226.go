package controllers

import (
	"github.com/gin-gonic/gin"
)

func UploadFile(context *gin.Context) {
	file, _ := context.FormFile("file")
}

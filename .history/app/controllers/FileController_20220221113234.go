package controllers

import (
	"github.com/gin-gonic/gin"
)

func UploadFile(context *gin.Context) {
	file, error := context.FormFile("file")

}

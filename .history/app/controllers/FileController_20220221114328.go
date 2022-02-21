package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	filename := filepath.Base(file.Filename)
	c.SaveUploadedFile(file, filename)
	if err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}
	c.String(http.StatusOK, "File %s uploaded successfully", file.Filename)
}

func UploadFiles(c *gin.Context) {
	files, err := c.FormFile("files")

}

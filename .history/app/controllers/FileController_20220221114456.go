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
	form, _ := c.MultipartForm()
	files := form.File["files"]
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}
	}
	c.String(http.StatusOK, "Uploaded successfully", len(files), name, email)
}

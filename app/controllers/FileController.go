/*
 * @title: 这里写标题
 * @version: 1.0
 * @author: huang sn
 * @description: 这里写描述信息
 * @BasePath: 这里写base path
 */
package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

/**
 * @description:
 * @param {*gin.Context} c
 * @return {*}
 */
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	// TODO:指定上传路径还没做
	filename := filepath.Base(file.Filename)
	c.SaveUploadedFile(file, filename)
	if err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}
	c.String(http.StatusOK, "File %s uploaded successfully", file.Filename)
}

/**
 * @description:
 * @param {*gin.Context} c
 * @return {*}
 */
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
	c.String(http.StatusOK, "Uploaded successfully")
}

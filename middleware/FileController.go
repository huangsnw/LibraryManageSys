package middleware

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
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

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
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

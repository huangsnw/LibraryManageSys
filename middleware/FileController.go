package middleware

import (
	"LibraryManageSys/pkg/pic"
	"LibraryManageSys/pkg/result"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @Summary 单个文件上传
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /file/upload [post]
func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	filename := strconv.Itoa(int(time.Now().UnixMicro())) + filepath.Base(file.Filename)
	// 当日时间
	time1 := time.Now().Format("2006-01-01")
	imagePath := viper.GetString("file.path") + "/" + time1 + "/"
	err1 := os.Mkdir(imagePath, os.ModePerm)
	if os.IsExist(err1) {
		log.Println("文件夹已经存在")
	}
	ctx.SaveUploadedFile(file, imagePath+filename)
	res := result.Result{}
	if err != nil {
		ctx.JSON(500, res.Error("文件上传失败"))
	} else {
		p := pic.Pic{
			Name: filename,
			Url:  viper.GetString("file.Url") + "/" + time1 + "/" + filename,
		}
		ctx.JSON(200, res.SUCESS(&p))
	}
}

// @Summary 多个文件上传
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /file/uploads [post]
func UploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	res := result.Result{}
	var filename string
	time1 := time.Now().Format("2006-01-01")
	imagePath := viper.GetString("file.path") + "/" + time1 + "/"
	for _, file := range files {
		filename = strconv.Itoa(int(time.Now().UnixMicro())) + filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, imagePath+filename); err != nil {
			c.JSON(500, res.Error("文件上传失败"))
			return
		}
	}
	p := pic.Pic{
		Name: filename,
		Url:  viper.GetString("file.Url") + "/" + time1 + "/" + filename,
	}
	c.JSON(200, res.SUCESS(&p))
}

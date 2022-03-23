package middleware

import (
	"LibraryManageSys/pkg/result"
	"LibraryManageSys/pkg/util"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	url := "/Users/huangsn/Pictures/一见倾心系列五周年4K壁纸（600张）/一见倾心系列五周年4K壁纸/汉服.国风"
	numbers, err := strconv.Atoi(ctx.Query("n"))
	folder, err := os.Open(url)
	util.CopeError(err, "文件打开出错1")
	files, err := folder.Readdirnames(numbers)
	util.CopeError(err, "文件打开出错2")
	var data = make([]string, numbers)
	for i, p := range files {
		data[i] = url + "/" + p
	}
	res := result.Result{}
	ctx.JSON(http.StatusOK, res.SUCESS(data))
}

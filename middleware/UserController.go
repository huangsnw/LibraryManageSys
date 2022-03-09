package middleware

import (
	"LibraryManageSys/models"
	"LibraryManageSys/pkg/page"
	"LibraryManageSys/pkg/result"
	"LibraryManageSys/pkg/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xuri/excelize/v2"
)

// @Summary 保存用户
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/save [post]
func SaveUser(ctx *gin.Context) {
	var user models.User
	ctx.ShouldBind(&user)
	util.DB.AutoMigrate(&models.User{})
	util.DB = util.DB.Create(&user)
	res := result.Result{}
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, res.OK())
	} else {
		ctx.JSON(500, res.ERROR())
	}
}

// @Summary 删除用户
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/del [post]
func DeleteUser(ctx *gin.Context) {
	var user models.User
	userId := ctx.Query("id")
	util.DB = util.DB.Delete(&user, "id = ? ", userId)
	res := result.Result{}
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, res.OK())
	} else {
		ctx.JSON(500, res.ERROR())
	}
}

// @Summary 选择用户
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/select [post]
func SelectUser(ctx *gin.Context) {
	var user models.User
	userId := ctx.Query("id")
	util.DB = util.DB.Take(&user, "id = ?", userId)
	res := result.Result{}
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, res.OK())
	} else {
		ctx.JSON(500, res.ERROR())
	}
}

// @Summary 更新用户
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/update [post]
func UpdateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	util.DB = util.DB.Model(&user).Updates(&user)
	res := result.Result{}
	if util.DB.RowsAffected == 1 {
		ctx.JSON(200, res.OK())
	} else {
		ctx.JSON(500, res.ERROR())
	}
}

// @Summary 导出用户
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/export [post]
func ExportUser(ctx *gin.Context) {
	var user []models.User
	util.DB = util.DB.Find(&user)

	f := excelize.NewFile()
	// 创建一个工作表
	index := f.NewSheet("Sheet1")
	// 设置标题
	f.SetCellValue("Sheet1", "A1", "姓名")
	f.SetCellValue("Sheet1", "B1", "身份证")
	f.SetCellValue("Sheet1", "C1", "手机号码")
	f.SetCellValue("Sheet1", "D1", "地址")
	f.SetCellValue("Sheet1", "E1", "等级")
	// 设置值
	nu := [...]string{"A", "B", "C", "D", "E"}
	for i, p := range user {
		j := strconv.Itoa(i + 2)
		f.SetCellValue("Sheet1", nu[i]+j, p.Name)
		f.SetCellValue("Sheet1", nu[i]+j, p.IdCard)
		f.SetCellValue("Sheet1", nu[i]+j, p.Phone)
		f.SetCellValue("Sheet1", nu[i]+j, p.Address)
		f.SetCellValue("Sheet1", nu[i]+j, p.Rank)
	}

	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	ctx.File("Book1.xlsx")
}

// @Summary 用户分页查询
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/page [get]
func PageSelectUser(ctx *gin.Context) {
	var sum int64
	var user []models.User
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", viper.GetString("page.pageSize")))
	pageNumber, _ := strconv.Atoi(ctx.DefaultQuery("pageNumber", viper.GetString("page.pageNumber")))
	util.DB.Model(&user).Count(&sum)
	err := util.DB.Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&user).Order("id desc").Error
	page := page.Page{}
	if err != nil {
		res := result.Result{}
		ctx.JSON(http.StatusInternalServerError, res.ERROR())
		return
	}
	ctx.JSON(http.StatusOK, page.OK(pageSize, pageNumber, int(sum), &user))
}

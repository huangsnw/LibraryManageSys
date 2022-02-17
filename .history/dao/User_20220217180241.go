package dao

import "time"

type User struct {
	// 主键ID
	ID string `json:"id" gorm:"id"`

	// 姓名
	Name string `json:"name" grom:"name"`

	// 身份证号码
	IdCard string

	// 电话
	Phone string

	// 居住地址
	Address string

	// 等级 (1: 普通用户    2：管理员用户)
	Rank int8

	// 创建时间
	CreateTime time.Time

	// 更新时间
	UpdateTime time.Time

	// 创建者
	CreateBy string

	// 更新者
	UpdateBy string
}

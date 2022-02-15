package dao

import "time"

type UserDao struct {
	// 主键ID
	Id string

	// 姓名
	Name string

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
}

package dao

import "time"

type BookDao struct {
	// 主键ID
	Id string

	// 书名
	BookName string

	// 作者
	Author string

	// 出版社
	Publisher string

	// ISBN
	Isbn string

	// 图书分类
	Classification string

	// 所属楼层
	Floor int8

	// 书架号码
	Bookshelf int8

	// 购买日期
	DateOfPurchase time.Time

	// 备注
	Remarks string

	// 创建时间
	CreateTime time.Time

	// 更新时间
	UpdateTime time.Time

	// 创建者
	CreateBy string

	// 更新者
	UpdateBy string
}

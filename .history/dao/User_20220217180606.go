package dao

import "time"

type User struct {
	ID      string `json:"id" gorm:"id"`
	Name    string `json:"name" gorm:"name"`
	IdCard  string `json:"id_card" gorm:"id_card"`
	Phone   string `json:"phone" gorm:"phone"`
	Address string `json:"address" gorm:"address"`
	Rank    int8   `json:"rank" gorm:"rank"`

	// 创建时间
	CreateTime time.Time

	// 更新时间
	UpdateTime time.Time

	// 创建者
	CreateBy string

	// 更新者
	UpdateBy string
}

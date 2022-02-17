package dao

import "time"

type User struct {
	ID         string    `json:"id" gorm:"id"`
	Name       string    `json:"name" gorm:"name"`
	IdCard     string    `json:"id_card" gorm:"id_card"`
	Phone      string    `json:"phone" gorm:"phone"`
	Address    string    `json:"address" gorm:"address"`
	Rank       int8      `json:"rank" gorm:"rank"`
	CreateAT   time.Time `json:"create_at" gorm:"create_at"`
	UpdateTime time.Time

	// 创建者
	CreateBy string

	// 更新者
	UpdateBy string
}

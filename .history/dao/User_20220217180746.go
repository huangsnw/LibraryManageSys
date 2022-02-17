package dao

import "time"

type User struct {
	ID       string    `json:"id" gorm:"id"`
	Name     string    `json:"name" gorm:"name"`
	IdCard   string    `json:"id_card" gorm:"id_card"`
	Phone    string    `json:"phone" gorm:"phone"`
	Address  string    `json:"address" gorm:"address"`
	Rank     int8      `json:"rank" gorm:"rank"`
	CreateAt time.Time `json:"create_at" gorm:"create_at"`
	UpdateAt time.Time `json:"update_at" gorm:"update_at"`
	CreateBy string    `json:"create_by" gorm:"create_by"`

	// 更新者
	UpdateBy string
}

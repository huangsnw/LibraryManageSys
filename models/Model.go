package models

import "time"

type Model struct {
	ID       string    `json:"id" gorm:"primaryKey" type:"uuid"`
	CreateAt time.Time `json:"create_at" gorm:"type:datetime"`
	UpdateAt time.Time `json:"update_at" gorm:"type:datetime"`
	CreateBy string    `json:"create_by" gorm:"type:varchar(32)"`
	UpdateBy string    `json:"update_by" gorm:"type:varchar(32)"`
}

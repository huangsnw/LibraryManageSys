package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Model   Model  `gorm:"embedded"`
	Name    string `json:"name" gorm:"name"`
	IdCard  string `json:"id_card" gorm:"id_card"`
	Phone   string `json:"phone" gorm:"phone"`
	Address string `json:"address" gorm:"address"`
	Rank    int8   `json:"rank" gorm:"rank"`
	Picture string `json:"picture" gorm:"column:picture;default:null"`
}

func (u *User) BeforeCreate(g *gorm.DB) (e error) {
	u.Model.ID = uuid.New()
	return
}

func (b *User) BeforeUpdate(g *gorm.DB) (e error) {
	b.Model.UpdatedAt = time.Now()
	return
}

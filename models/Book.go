package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Model
	BookName       string     `json:"book_name" gorm:"type:varchar(32)"`
	Author         string     `json:"author" gorm:"type:varchar(32)"`
	Publisher      string     `json:"publisher" gorm:"type:varchar(32)"`
	Isbn           string     `json:"isbn" gorm:"type:varchar(32)"`
	Classification string     `json:"classification" gorm:"type:varchar(32)"`
	Floor          int8       `json:"floor" gorm:"type:int(11)"`
	Bookshelf      int8       `json:"bookshelf" gorm:"type:int(11)"`
	DateOfPurchase *time.Time `json:"date_of_purchase" gorm:"default:null"`
	Remarks        string     `json:"remarks" gorm:"type:varchar(32)"`
	Picture        string     `json:"picture" gorm:"column:picture;type:varchar(255)"`
}

func (b *Book) BeforeCreate(g *gorm.DB) (e error) {
	b.Model.ID = uuid.New()
	return
}

func (b *Book) BeforeUpdate(g *gorm.DB) (e error) {
	b.Model.UpdatedAt = time.Now()
	return
}

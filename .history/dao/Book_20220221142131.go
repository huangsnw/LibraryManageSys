package dao

import "time"

type Book struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	BookName       string    `json:"book_name" gorm:"type:varchar(32)"`
	Author         string    `json:"author" gorm:"type:varchar(32)"`
	Publisher      string    `json:"publisher" gorm:"type:varchar(32)"`
	Isbn           string    `json:"isbn" gorm:"type:varchar(32)"`
	Classification string    `json:"classification" gorm:"type:varchar(32)"`
	Floor          int8      `json:"floor" gorm:"type:int(11)"`
	Bookshelf      int8      `json:"bookshelf" gorm:"type:int(11)"`
	DateOfPurchase time.Time `json:"date_of_purchase" gorm:"type:timestamp"`
	Remarks        string    `json:"remarks" gorm:"type:varchar(32)"`
	CreateAt       time.Time `json:"create_at" gorm:"type:timestamp"`
	UpdateAt       time.Time `json:"update_at" gorm:"type:timestamp"`
	CreateBy       string    `json:"create_by" gorm:"type:varchar(32)"`
	UpdateBy       string    `json:"update_by" gorm:"type:varchar(32)"`
}

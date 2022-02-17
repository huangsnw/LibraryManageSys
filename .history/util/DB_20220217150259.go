package util

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:Root123.@tcp(1.15.189.106)/library"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error")
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

package util

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	database := viper.GetString("datasource.database")
	dsn := username + ":" + password + "@tcp(" + host + ")/" + database + "?parseTime=true"
	log.Println("Database Connect Sucess")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error")
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

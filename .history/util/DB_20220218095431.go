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
	password := viper.GetString("datasource.passowrd")
	host := viper.GetString("datasource.host")
	database:=viper.GetString("datasource.database")
	// dsn := "root:Root123.@tcp(1.15.189.106)/library"
	dsn:=
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error")
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

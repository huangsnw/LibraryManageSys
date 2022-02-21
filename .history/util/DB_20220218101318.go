package util

import (
	"LibraryManageSys/storage/logs"
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
	database := viper.GetString("datasource.database")
	// dsn := "root:Root123.@tcp(1.15.189.106)/library"
	dsn := username + ":" + password + "@tcp(" + host + ")/" + database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error")
	}
	DB = db
	logs.Log.Infow("数据库连接信息:\n")
	logs.Log.Infow(Host, host)
}

func GetDB() *gorm.DB {
	return DB
}
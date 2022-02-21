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
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	database := viper.GetString("datasource.database")
	// dsn := "root:Root123.@tcp(1.15.189.106)/library"
	dsn := username + ":" + password + "@tcp(" + host + ")/" + database
	logs.Log.Info("数据库连接信息:\n")
	logs.Log.Info("Host:", host)
	logs.Log.Info("Database:", database)
	logs.Log.Info("UserName:", username)
	logs.Log.Info("Password:", password)
	logs.Log.Info("dsn:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error")
	}
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

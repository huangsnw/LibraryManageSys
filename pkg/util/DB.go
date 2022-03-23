package util

import (
	"log"
	"time"

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
	dsn := username + ":" + password + "@tcp(" + host + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	log.Println("Database Connect Sucess")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})
	if err != nil {
		log.Fatal("Error")
	}
	// 链接池配置
	// sqlDB, err := db.DB()
	// 设置空闲连接池中连接的最大数量
	// sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	// sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间
	// sqlDB.SetConnMaxLifetime(time.Hour)
	// if err != nil {
	// 	log.Fatal("Error")
	// }
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}

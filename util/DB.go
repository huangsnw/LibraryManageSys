package util

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/url"
)

var DB *gorm.DB

func Init() {
	// driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	print(host)
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	fmt.Printf(args)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库链接出错")
	}
	DB = db
}

/*
 * @title: 这里写标题
 * @Date: 2022-02-16 10:48:23
 * @version: 1.0
 * @author: huang sn
 * @description: 这里写描述信息
 * @FilePath: /LibraryManageSys/util/DB.go
 */
package util

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

/**
 * @description:
 * @Accept:
 * @param {*}
 * @return {*}
 * @Router:
 */
func InitDB() {
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	database := viper.GetString("datasource.database")
	dsn := username + ":" + password + "@tcp(" + host + ")/" + database
	log.Println("Database Connect Sucess")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error")
	}
	DB = db
}

/**
 * @description:
 * @Accept:
 * @param {*}
 * @return {*}
 * @Router:
 */
func GetDB() *gorm.DB {
	return DB
}

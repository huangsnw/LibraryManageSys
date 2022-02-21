package util

import (
	"github.com/spf13/viper"
	"log"
)

func ConfigViper() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config-dev")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Viper解析出错")
	}
}

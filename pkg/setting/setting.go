package setting

import (
	"log"

	"github.com/spf13/viper"
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

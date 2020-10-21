package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func ReadConfig(){
	viper.SetConfigName("config.dev")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Print("Can't read config from file config")
	}

	logrus.Print("Config file read successfully, enjoy!")
}
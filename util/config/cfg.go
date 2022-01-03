package config

import (
	"github.com/riceChuang/gamerobot/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configuration *model.ServerConfig

// initial configs
func LoadConfig() *model.ServerConfig {
	configuration = &model.ServerConfig{}
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading configs file, %s", err)
	}
	viper.SetConfigName("game_list")
	viper.AddConfigPath("./configs/")
	viper.SetConfigType("yaml")
	viper.MergeInConfig()
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Panicf("Unable to decode configs into struct, %v", err)
	}

	return configuration
}

func GetInstance() *model.ServerConfig {
	return configuration
}

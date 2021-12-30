package config

import (
	"gitlab.baifu-tech.net/dsg-game/game-robot/model"
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

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Panicf("Unable to decode configs into struct, %v", err)
	}

	return configuration
}

func GetInstance() *model.ServerConfig {
	return configuration
}

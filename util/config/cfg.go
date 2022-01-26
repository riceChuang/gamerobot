package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/riceChuang/gamerobot/model"
	"sync"
)

var configuration *model.ServerConfig
var gameConfig map[string]*model.RoomCfgBase
var mu sync.Mutex

// initial config
func LoadConfig() *model.ServerConfig {
	configuration = &model.ServerConfig{}
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}
	viper.SetConfigName("game_list")
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yaml")
	viper.MergeInConfig()
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Panicf("Unable to decode config into struct, %v", err)
	}
	fmt.Println(configuration)
	return configuration
}

func LoadGameConfig() map[string]*model.RoomCfgBase {
	gameConfig = make(map[string]*model.RoomCfgBase)
	gameViper := viper.New()
	gameViper.SetConfigType("yaml")

	for _, game := range configuration.GameList {
		gameViper.SetConfigName(fmt.Sprintf("%v.robotConfig", game.GameName))
		gameViper.AddConfigPath(fmt.Sprintf("./config/robotconfig/%v", game.GameName))
		if err := gameViper.ReadInConfig(); err != nil {
			log.Panicf("Error reading config file, %s", err)
		}
		var roomCfg = &model.RoomCfg{}
		err := gameViper.Unmarshal(&roomCfg)
		if err != nil {
			log.Panicf("Unable to decode config into struct, %v", err)
		}
		for _, room := range roomCfg.Rooms {
			gameConfig[fmt.Sprintf("%v-%v", game.GameID, room.RoomIndex)] = room
		}
	}

	return gameConfig
}

func GetInstance() *model.ServerConfig {
	return configuration
}

func GetGameInstance() map[string]*model.RoomCfgBase {
	return gameConfig
}

func GetGameInstanceByID(id string) *model.RoomCfgBase {
	//id組合方法 GameID-RoomIndex e.x. 30-1
	mu.Lock()
	defer mu.Unlock()
	return gameConfig[id]
}

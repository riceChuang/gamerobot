package model

type ServerConfig struct {
	BindPort        string      `mapstructure:"bindport"`
	LogLevel        string      `mapstructure:"loglevel"`
	TestUserAccount string      `mapstructure:"account"`
	TestUserPwd     string      `mapstructure:"password"`
	Domain          string      `mapstructure:"domain"`
	Environment     []EnvCfg    `mapstructure:"environment"`
	GameList        []CommonCfg `mapstructure:"games"`
}

// GameEnvCfg ...
type EnvCfg struct {
	ENV         string `mapstructure:"env"`
	ServerURL   string `mapstructure:"serverUrl"`
	LoginDomain string `mapstructure:"loginDomain"`
	AgentID     int    `mapstructure:"agentID"`
}

// CommonCfg ...
type CommonCfg struct {
	RoomType []string `mapstructure:"roomtype"`
	GameID   int32    `mapstructure:"gameid"`
	HallPort int32    `mapstructure:"hallport"`
	GameName string   `mapstructure:"name"`
}

// RoomCfgBase ...
type RoomCfgBase struct {
	RoomIndex   int32                  `mapstructure:"roomIndex"`
	RobotNumber int32                  `mapstructure:"robotNumber"`
	MinMoney    int32                  `mapstructure:"minMoney"`
	StoredMoney int32                  `mapstructure:"storedMoney"`
	RobotName   string                 `mapstructure:"robotName"`
	Strategy    map[string]interface{} `mapstructure:"strategy"`
	RobotCount  int32
}

type RoomCfg struct {
	Name  string         `mapstructure:"name"`
	Rooms []*RoomCfgBase `mapstructure:"rooms"`
}

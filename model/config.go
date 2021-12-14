package model

type ServerConfig struct {
	BindPort        string      `mapstructure:"bindport"`
	LogLevel        string      `mapstructure:"loglevel"`
	TestUserAccount string      `mapstructure:"account"`
	TestUserPwd     string      `mapstructure:"password"`
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
}


// RoomCfgBase ...
type RoomCfgBase struct {
	RoomIndex   int32  `json:"roomIndex"`
	RobotNumber int32  `json:"robotNumber"`
	MinMoney    int32  `json:"minMoney"`
	StoredMoney int32  `json:"storedMoney"`
	RobotName   string `json:"robotName"`
}

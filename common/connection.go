package common

type DispatcherType string

const (
	ClientWSDispatcher DispatcherType = "client_ws_dispatcher"
	GameWSDispatcher   DispatcherType = "game_ws_dispatcher"
)

type PasserType string

const (
	Client               PasserType = "client"
	Game                 PasserType = "game"
	ClientServerTransfer PasserType = "client_server_transfer"
	GameServerTransfer   PasserType = "game_server_transfer"
	ClientSender         PasserType = "client_sender"
	GameSender           PasserType = "game_sender"
	ClientFileWriter     PasserType = "client_to_file"
)

type ConnectType int32

const (
	HttpConnect ConnectType = 1
	GameConnect ConnectType = 2
)

const (
	WSClose = 99
)

const (
	SendTypeBoth             = 1       //送client+game robot
	SendGameOnly             = 2        //送game robot
	SendTypeClientWithFilter = 4       //送client 讀取config底下 需要的id
	//FileWriter = 5       //把client 的file寫出
)

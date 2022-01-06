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
)

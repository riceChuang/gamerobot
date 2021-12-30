package common

type DispatcherType string

const (
	ClientWSDispatcher DispatcherType = "client_ws_dispatcher"
	GameWSDispatcher   DispatcherType = "game_ws_dispatcher"
)

type PasserType string

const (
	ClientToServer   PasserType = "client_to_server"
	GameToServer     PasserType = "game_to_server"
	ServerToTransfer PasserType = "server_to_transfer"
	TransferToClient PasserType = "transfer_to_client"
	TransferToGame   PasserType = "transfer_to_game"
)

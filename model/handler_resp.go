package model

type IndexResp struct {
	GameList map[int32]*GameInfo
	Envs     []string
}

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type LoginRespData struct {
	ConnID   string `json:"conn"`
	Account  string `json:"account"`
	Env      string `json:"env"`
	GameName string `json:"gamename"`
}

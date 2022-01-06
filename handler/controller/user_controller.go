package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/framework/gametype"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/service/connect"
	"github.com/riceChuang/gamerobot/usecase"
	"github.com/riceChuang/gamerobot/util/config"
	"github.com/riceChuang/gamerobot/util/logs"
	log "github.com/sirupsen/logrus"

	"net/http"
)

type UserController struct {
	Config  *model.ServerConfig
	useCase *usecase.UseCase
	logger  *log.Entry
}

func NewUserController() *UserController {
	return &UserController{
		Config:  config.GetInstance(),
		useCase: usecase.NewUseCase(),
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "user_controller",
		}),
	}
}

//取得root
func (h *UserController) GetIndex(ctx *gin.Context) {

	resp := model.IndexResp{
		GameList: make(map[int32]*model.GameInfo),
		Envs:     []string{},
	}
	//取得遊戲列表
	for _ , game := range h.Config.GameList {
		instance := gametype.GetInstanceByContentType(common.GameServerID(game.GameID))
		if instance == nil {
			continue
		}
		resp.GameList[game.GameID] = &model.GameInfo{
			GameID:   game.GameID,
			Name:     common.GameServerIDToString(common.GameServerID(game.GameID)),
			Buttons:  instance.GetGameHandler().GetMessageBtn(),
			RoomType: game.RoomType,
		}
	}

	//取得環境列表
	for _, value := range h.Config.Environment {
		resp.Envs = append(resp.Envs, value.ENV)
	}

	ctx.HTML(http.StatusOK, "views/game_cascade.html", gin.H{
		"resp": resp,
	})
}

//用戶登入
func (h *UserController) UserLogin(ctx *gin.Context) {
	request := &model.UserLogin{}

	err := ctx.BindJSON(request)
	if err != nil {
		h.logger.Error(err)
		return
	}
	if request.UserName == "" {
		h.logger.Debug("沒設定name 使用default")
		request.UserName = h.Config.TestUserAccount
	}

	if request.Password == "" {
		h.logger.Debug("沒設定password 使用default")
		request.Password = h.Config.TestUserPwd
	}

	if request.GameID == 0 {
		h.logger.Info("沒有設定遊戲Id")
		ctx.Error(model.ErrInvalidRequest)
		return
	}

	if request.Env == "" {
		h.logger.Info("沒有設定環境Id")
		ctx.Error(model.ErrInvalidRequest)
		return
	}

	if request.RoomType == "" {
		h.logger.Info("沒有設定房間index")
		ctx.Error(model.ErrInvalidRequest)
		return
	}
	gameInstance := gametype.GetInstanceByContentType(common.GameServerID(request.GameID))
	gConfig := gameInstance.GetGameConfig()
	var gameRoom string
	for flag, room := range gConfig.RoomType {
		if room == request.RoomType {
			gameRoom = fmt.Sprintf("%v-%v", gConfig.GameID, flag+1)
			break
		}
	}

	var envInfo *model.EnvCfg
	for _, env := range h.Config.Environment {
		if env.ENV == request.Env {
			envInfo = &env
			break
		}
	}
	if envInfo == nil {
		h.logger.Error("找不到環境")
		ctx.Error(model.ErrInternalServerError)
		return
	}

	dsgLoginReq := &model.DSGLoginReq{
		LoginDomain: envInfo.LoginDomain,
		AgentID:     envInfo.AgentID,
		Account:     request.UserName,
		Password:    request.Password,
		GameID:      request.GameID,
	}
	//取得token
	var dsgLoginResp *model.DSGLoginResp
	dsgLoginResp, err = h.useCase.DsgAPIBase.Login(dsgLoginReq)
	if err != nil {
		h.logger.Error(err)
		ctx.Error(model.ErrInternalServerError)
		return
	}
	h.logger.Info(dsgLoginResp.URL)
	//取得token 後 用token取得向hall大廳取得遊戲port號
	hallURL := fmt.Sprintf("%s:%d", envInfo.ServerURL, gConfig.HallPort)
	gamePort, userInfo := h.useCase.AdminBase.GetGameWsInfo(hallURL, gameRoom, dsgLoginResp.UserName, dsgLoginResp.Token)

	//取得遊戲port號後 新增game connect
	gameURL := fmt.Sprintf("%s:%s", envInfo.ServerURL, gamePort)
	gameConnect := connect.NewGameConnect(framework.NewProtoConnect(framework.NewConn(gameURL, nil)), userInfo)

	//新增client conn
	conn, err := connect.NewClientWsGameConn(ctx)
	if err != nil {
		h.logger.Error("conn error: %v", err)
		return
	}
	conn.SetGameConn(gameConnect)
	conn.SetGameID(common.GameServerID(request.GameID))

	//將client conn 加入到manager中
	connManager := connect.GetClientGameCommunicateManager()
	connManager.AddClient(conn)

	ctx.JSON(http.StatusOK, model.Resp{
		Code:    http.StatusOK,
		Message: "success",
		Data:    conn.ID,
	})
	return
}

//用戶發送訊息
func (h *UserController) SendMessage(ctx *gin.Context) {
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *UserController) WebSocketConn(ctx *gin.Context) {
	connID := ctx.Query("connID")
	fmt.Println(connID)

	connManager := connect.GetClientGameCommunicateManager()
	conn := connManager.GetClient(connID)
	if conn == nil {
		h.logger.Errorf("找不到connID:%v", connID)
		ctx.Error(model.ErrNotFound)
		return
	}

	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	connection := framework.NewConn("", ws)
	wsConnection := framework.NewHttpConnect(connection)
	conn.SetWsConn(wsConnection)

	err = wsConnection.Connect()
	if err != nil {
		h.logger.Error(err)
		return
	}

	//連線gameWS
	err = conn.ProtoConnect()
	if err != nil {
		h.logger.Error(err)
		return
	}
	return
}

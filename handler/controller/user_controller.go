package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/riceChuang/gamerobot/common"
	connection "github.com/riceChuang/gamerobot/framework/connection/connect"
	"github.com/riceChuang/gamerobot/framework/connection/connecttype"
	"github.com/riceChuang/gamerobot/game/gametype"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/service/connect"
	"github.com/riceChuang/gamerobot/usecase"
	"github.com/riceChuang/gamerobot/util/config"
	"github.com/riceChuang/gamerobot/util/logs"

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
func (h *UserController) GetIndexPage(ctx *gin.Context) {

	resp := model.IndexResp{
		GameList: make(map[int32]*model.GameInfo),
		Envs:     []string{},
	}
	//取得遊戲列表
	for _, game := range h.Config.GameList {
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
	resp.Domain = h.Config.Domain
	ctx.HTML(http.StatusOK, "views/game_cascade.html", gin.H{
		"resp": resp,
	})
}

//取得壓測頁面
func (h *UserController) GetStressPage(ctx *gin.Context) {
	resp := model.IndexResp{
		GameList: make(map[int32]*model.GameInfo),
		Envs:     []string{},
	}
	//取得遊戲列表
	for _, game := range h.Config.GameList {
		instance := gametype.GetInstanceByContentType(common.GameServerID(game.GameID))
		if instance == nil {
			continue
		}
		resp.GameList[game.GameID] = &model.GameInfo{
			GameID:       game.GameID,
			Name:         common.GameServerIDToString(common.GameServerID(game.GameID)),
			Buttons:      instance.GetGameHandler().GetMessageBtn(),
			RoomType:     game.RoomType,
			RoomStrategy: make(map[string]map[string]interface{}),
		}
		for roomIndex, t := range game.RoomType {
			roomCfg := config.GetGameInstanceByID(fmt.Sprintf("%v-%v", game.GameID, roomIndex+1))
			if roomCfg != nil {
				resp.GameList[game.GameID].RoomStrategy[t] = roomCfg.Strategy
			}
		}
	}

	//取得環境列表
	for _, value := range h.Config.Environment {
		resp.Envs = append(resp.Envs, value.ENV)
	}
	resp.Domain = h.Config.Domain
	ctx.HTML(http.StatusOK, "views/robot_test.html", gin.H{
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest,model.ErrInvalidRequest)
		return
	}

	if request.Env == "" {
		h.logger.Info("沒有設定環境Id")
		ctx.AbortWithStatusJSON(http.StatusBadRequest,model.ErrInvalidRequest)
		return
	}

	if request.RoomType == "" {
		h.logger.Info("沒有設定房間index")
		ctx.AbortWithStatusJSON(http.StatusBadRequest,model.ErrInvalidRequest)
		return
	}

	var roomIndex int
	for _, gameCfg := range h.Config.GameList {
		if gameCfg.GameID == request.GameID {
			for index, name := range gameCfg.RoomType {
				if name == request.RoomType {
					roomIndex = index + 1
					break
				}
			}
			break
		}
	}

	gameInstance := gametype.GetInstanceByContentType(common.GameServerID(request.GameID), int32(roomIndex))
	gConfig := gameInstance.GetGameConfig()

	roomCfg := config.GetGameInstanceByID(gameInstance.GetGameHandler().GetGameRoomIndex())
	if roomCfg == nil {
		h.logger.Info("找不到cofig")
		ctx.AbortWithStatusJSON(http.StatusBadRequest,model.APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "config not found"})
		return
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest,model.ErrInternalServerError)
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest,model.ErrInternalServerError)
		return
	}
	h.logger.Info(dsgLoginResp.URL)

	dsgStoreReq := &model.DSGStoreMoneyReq{
		LoginDomain: envInfo.LoginDomain,
		AgentID:     envInfo.AgentID,
		Account:     request.UserName,
		Money:       roomCfg.StoredMoney,
	}
	//取得token
	var dsgStoreResp *model.DSGStoreMoneyResp
	dsgStoreResp, err = h.useCase.DsgAPIBase.StoreMoney(dsgStoreReq)
	if err != nil {
		h.logger.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,model.ErrInternalServerError)
		return
	}
	h.logger.Info(dsgStoreResp.Money)

	//取得token 後 用token取得向hall大廳取得遊戲port號
	hallURL := fmt.Sprintf("%s:%d", envInfo.ServerURL, gConfig.HallPort)
	gamePort, userInfo := h.useCase.AdminBase.GetGameWsInfo(hallURL, gameInstance.GetGameHandler().GetGameRoomIndex(), dsgLoginResp.UserName, dsgLoginResp.Token)

	//取得遊戲port號後 新增game connect
	gameURL := fmt.Sprintf("%s:%s", envInfo.ServerURL, gamePort)
	gameConnect := connecttype.NewGameConnect(envInfo, connection.NewProtoConnect(connection.NewConn(gameURL, nil, common.GameConnect)), userInfo)

	//新增client conn
	conn, err := connect.NewClientWsGameConn(ctx)
	if err != nil {
		h.logger.Error("conn error: %v", err)
		return
	}
	conn.SetGameConn(gameConnect)
	conn.SetGameType(gameInstance)
	conn.SetGameID(common.GameServerID(request.GameID))

	//將client conn 加入到manager中
	connManager := connect.GetClientGameCommunicateManager()
	connManager.AddClient(conn)

	ctx.JSON(http.StatusOK, model.Resp{
		Code:    http.StatusOK,
		Message: "success",
		Data: model.LoginRespData{
			ConnID:   conn.ID,
			Account:  request.UserName,
			Env:      request.Env,
			GameName: fmt.Sprintf("%v-%v", gameInstance.GetGameConfig().GameName, request.RoomType),
		},
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
	h.logger.Infof("myconnID:%v", connID)

	connManager := connect.GetClientGameCommunicateManager()
	conn := connManager.GetClient(connID)
	if conn == nil {
		h.logger.Errorf("找不到connID:%v", connID)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,model.ErrNotFound)
		return
	}

	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		h.logger.Error(err)
		return
	}
	httpConnect := connection.NewConn("", ws, common.HttpConnect)
	wsConnection := connection.NewHttpConnect(httpConnect)
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

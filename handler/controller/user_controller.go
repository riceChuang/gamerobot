package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/riceChuang/gamerobot/common"
	"github.com/riceChuang/gamerobot/framework"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/service/connect"
	"github.com/riceChuang/gamerobot/service/game"
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
func (h *UserController) GetIndex(ctx *gin.Context) {

	resp := model.IndexResp{
		GameList: make(map[int32]*model.GameInfo),
		Envs:     []string{},
	}

	//取得遊戲列表
	allGames := game.GetAllGameInstance()
	for game, instance := range allGames {
		resp.GameList[int32(game)] = &model.GameInfo{
			GameID:  instance.GameID(),
			Name:    instance.GameName(),
			Buttons: instance.GetMessageBtn(),
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
	var token string
	token, err = h.useCase.DsgAPIBase.Login(envInfo.LoginDomain, envInfo.AgentID, request.UserName, request.Password, request.GameID)
	if err != nil {
		h.logger.Error(err)
		ctx.Error(model.ErrInternalServerError)
		return
	}
	h.logger.Info(token)

	connManager := connect.GetClientGameCommunicateManager()
	conn, err := connect.NewClientWsGameConn(ctx)
	if err != nil {
		h.logger.Error("conn error: %v", err)
		return
	}
	connManager.AddClient(conn)
	conn.SetGameConn(connect.NewProtoConnect(framework.NewConn(token, nil)))
	conn.SetGameID(common.GameServerID(request.GameID))

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
	wsConnection := connect.NewHttpConnect(connection)
	conn.SetWsConn(wsConnection)

	err = wsConnection.Connect()
	if err != nil {
		h.logger.Error(err)
		return
	}


	return
}

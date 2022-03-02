package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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

type StressController struct {
	Config  *model.ServerConfig
	useCase *usecase.UseCase
	logger  *log.Entry
}

func NewStressController() *StressController {
	return &StressController{
		Config:  config.GetInstance(),
		useCase: usecase.NewUseCase(),
		logger: logs.GetLogger().WithFields(log.Fields{
			"server": "user_controller",
		}),
	}
}

//取得遊戲列表
func (sc *StressController) GetGameList(ctx *gin.Context) {

	resp := model.Resp{
		Code:    http.StatusOK,
		Message: "success",
	}

	gameListResp := model.GameListResp{
		GameList: make(map[int32]*model.GameInfo),
	}
	//取得遊戲列表
	for _, game := range sc.Config.GameList {
		instance := gametype.GetInstanceByContentType(common.GameServerID(game.GameID))
		if instance == nil {
			continue
		}
		gameListResp.GameList[game.GameID] = &model.GameInfo{
			GameID:   game.GameID,
			Name:     common.GameServerIDToString(common.GameServerID(game.GameID)),
			Buttons:  instance.GetGameHandler().GetMessageBtn(),
			RoomType: game.RoomType,
		}
	}
	resp.Data = gameListResp
	ctx.JSON(http.StatusOK, resp)

}

//啟動壓測機器人
func (sc *StressController) RunStress(ctx *gin.Context) {

	request := []*model.StressReq{}

	err := ctx.BindJSON(&request)
	if err != nil {
		sc.logger.Error(err)
		return
	}

	for _, stressReq := range request {
		if stressReq.GameID == 0 {
			sc.logger.Info("沒有設定遊戲Id")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "gameid is not set"})
			return
		}

		if stressReq.Env == "" {
			sc.logger.Info("沒有設定環境Id")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "env is not set"})
			return
		}

		if stressReq.RoomType == "" {
			sc.logger.Info("沒有設定房間index")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "roomtype is not set"})
			return
		}

		if stressReq.RobotCount <= 0 {
			sc.logger.Info("沒有設定機器人壓測數量")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "robotcount is not set"})
		}

		var envInfo *model.EnvCfg
		for _, env := range sc.Config.Environment {
			if env.ENV == stressReq.Env {
				envInfo = &env
				break
			}
		}

		if envInfo == nil {
			sc.logger.Error("找不到環境")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrInternalServerError)
			return
		}

		var strategy map[string]interface{}
		err = json.Unmarshal([]byte(request[0].Strategy), &strategy)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "設定Strategy Error"})
			return
		}

		var roomIndex int
		for _, gameCfg := range sc.Config.GameList {
			if gameCfg.GameID == stressReq.GameID {
				for index, name := range gameCfg.RoomType {
					if name == stressReq.RoomType {
						roomIndex = index + 1
						break
					}
				}
				break
			}
		}

		gameInstance := gametype.GetInstanceByContentType(common.GameServerID(stressReq.GameID), int32(roomIndex))
		err = gameInstance.GetGameHandler().SetStrategy(strategy)
		if err != nil {
			sc.logger.Error("設定Strategy Error")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "設定Strategy Error"})
			return
		}
		gConfig := gameInstance.GetGameConfig()

		roomCfg := config.GetGameInstanceByID(gameInstance.GetGameHandler().GetGameRoomIndex())
		if roomCfg == nil {
			sc.logger.Info("找不到cofig")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "config not found"})
			return
		}

		if int32(stressReq.RobotCount) > (roomCfg.RobotNumber - roomCfg.RobotCount) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.APIException{Code: http.StatusBadRequest, Message: "bad_request", Data: "超過可使用的機器人數量"})
			return
		}

		for i := 1; i <= stressReq.RobotCount; i++ {
			//設定機器人號碼
			roomCfg.RobotCount++
			robotCount := fmt.Sprintf("%04d", roomCfg.RobotCount)

			dsgLoginReq := &model.DSGLoginReq{
				LoginDomain: envInfo.LoginDomain,
				AgentID:     envInfo.AgentID,
				Account:     fmt.Sprintf("%v_%v", roomCfg.RobotName, robotCount),
				GameID:      stressReq.GameID,
			}

			//取得token
			var dsgLoginResp *model.DSGLoginResp
			dsgLoginResp, err = sc.useCase.DsgAPIBase.Login(dsgLoginReq)
			if err != nil {
				sc.logger.Error(err)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrInternalServerError)
				return
			}

			sc.logger.Info(dsgLoginResp.URL)

			dsgStoreReq := &model.DSGStoreMoneyReq{
				LoginDomain: envInfo.LoginDomain,
				AgentID:     envInfo.AgentID,
				Account:     fmt.Sprintf("%v_%v", roomCfg.RobotName, robotCount),
				Money:       roomCfg.StoredMoney,
			}
			//取得token
			var dsgStoreResp *model.DSGStoreMoneyResp
			dsgStoreResp, err = sc.useCase.DsgAPIBase.StoreMoney(dsgStoreReq)
			if err != nil {
				sc.logger.Error(err)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrInternalServerError)
				return
			}
			sc.logger.Info(dsgStoreResp.Money)

			go func() {
				//取得token 後 用token取得向hall大廳取得遊戲port號
				hallURL := fmt.Sprintf("%s:%d", envInfo.ServerURL, gConfig.HallPort)
				gamePort, userInfo := sc.useCase.AdminBase.GetGameWsInfo(hallURL, gameInstance.GetGameHandler().GetGameRoomIndex(), dsgLoginResp.UserName, dsgLoginResp.Token)

				//取得遊戲port號後 新增game connect
				gameURL := fmt.Sprintf("%s:%s", envInfo.ServerURL, gamePort)
				gameConnect := connecttype.NewGameConnect(envInfo, connection.NewProtoConnect(connection.NewConn(gameURL, nil, common.GameConnect)), userInfo)

				//新增client conn
				conn, err := connect.NewClientWsGameConn(ctx)
				if err != nil {
					sc.logger.Error("conn error: %v", err)
					return
				}
				conn.SetGameConn(gameConnect)
				conn.SetGameType(gameInstance)
				conn.SetGameID(common.GameServerID(stressReq.GameID))

				//將client conn 加入到manager中
				connManager := connect.GetClientGameCommunicateManager()
				connManager.AddClient(conn)

				err = conn.ProtoConnect()
				if err != nil {
					sc.logger.Error("protoconn Error: %v", err)
					return
				}

			}()

		}
	}

	ctx.JSON(http.StatusOK, model.Resp{
		Code:    http.StatusOK,
		Message: "success",
		Data:    model.LoginRespData{},
	})
	return
}

func (sc *StressController) WebSocketConn(ctx *gin.Context) {
	connID := ctx.Query("connID")
	fmt.Println(connID)

	connManager := connect.GetClientGameCommunicateManager()
	allConns := connManager.GetAllClient()

	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	httpConnect := connection.NewConn("", ws, common.HttpConnect)
	wsConnection := connection.NewHttpConnect(httpConnect,true)

	for _, conn := range allConns {
		conn.SetWsConn(wsConnection)
	}

	err = wsConnection.Connect()
	if err != nil {
		sc.logger.Error(err)
		return
	}
	return
}

//取得壓測機器人狀態
func (sc *StressController) GetStatus(ctx *gin.Context) {

}

//清除所有壓測機器人
func (sc *StressController) ClearRobot(ctx *gin.Context) {
	connManager := connect.GetClientGameCommunicateManager()
	allClients := connManager.GetAllClient()
	if len(allClients) > 0 {
		for _, client := range allClients {
			err := client.CloseProtoConnect()
			if err != nil {
				sc.logger.Errorf("關閉遊戲ws錯誤:%v", err)
				continue
			}
			err = client.CloseHttpConnect()
			if err != nil {
				sc.logger.Errorf("關閉httpws錯誤:$v", err)
			}
		}
	}

	allGame := config.GetGameInstance()
	for _, gameCfg := range allGame {
		gameCfg.RobotCount = 0
	}
	ctx.Status(http.StatusOK)
	return
}

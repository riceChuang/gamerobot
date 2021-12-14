package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/riceChuang/gamerobot/game"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/service/usecase"
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

	ctx.HTML(http.StatusOK, "views/index.html", gin.H{
		"resp": resp,
	})
}

//用戶登入
func (h *UserController) UserLogin(ctx *gin.Context) {
	userLoginJson := &model.UserLogin{}

	ctx.BindJSON(userLoginJson)
	if userLoginJson.Name == "" {
		logs.GetLogger().Debug("沒設定name 使用default")
		userLoginJson.Name = h.Config.TestUserAccount
	}

	if userLoginJson.Password == "" {
		logs.GetLogger().Debug("沒設定password 使用default")
		userLoginJson.Password = h.Config.TestUserPwd
	}

	if userLoginJson.GameID == 0 {
		logs.GetLogger().Info("沒有設定遊戲Id")
		ctx.Error(model.ErrInvalidRequest)
		return
	}

	if userLoginJson.Env == 0 {
		logs.GetLogger().Info("沒有設定環境Id")
		ctx.Error(model.ErrInvalidRequest)
		return
	}


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
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	go func() {
		defer ws.Close()
		for {
			//读取ws中的数据
			mt, message, err := ws.ReadMessage()
			if err != nil {
				break
			}
			if string(message) == "ping" {
				message = []byte("pong")
			}
			//写入ws数据
			err = ws.WriteMessage(mt, message)
			if err != nil {
				break
			}
		}
	}()
}


func (h *UserController) ChatIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "chat_index.html", gin.H{})

}
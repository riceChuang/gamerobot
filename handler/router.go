package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/riceChuang/gamerobot/handler/controller"
)

func Router(router *gin.Engine) {

	controlManager := controller.NewController()
	//router.Handle(/ws,controlManager.User.WebSocketConn))

	router.GET("/", controlManager.User.GetIndexPage)
	router.GET("/stress", controlManager.User.GetStressPage)
	router.GET("/ws", controlManager.User.WebSocketConn)
	router.POST("/login", controlManager.User.UserLogin)
	router.POST("/sendmessage", controlManager.User.SendMessage)
	router.GET("/stressws", controlManager.Stress.WebSocketConn)
	router.GET("/gamelist", controlManager.Stress.GetGameList)
	router.GET("/status", controlManager.Stress.GetStatus)
	router.GET("/stopstress", controlManager.Stress.ClearRobot)
	router.POST("/startstress", controlManager.Stress.RunStress)

}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/riceChuang/gamerobot/handler/controller"
)

func Router(router *gin.Engine) {

	controlManager := controller.NewController()
	//router.Handle(/ws,controlManager.User.WebSocketConn))

	router.GET("/", controlManager.User.GetIndex)
	router.GET("/ws",controlManager.User.WebSocketConn)
	router.POST("/login", controlManager.User.UserLogin)
	router.POST("/sendmessage", controlManager.User.SendMessage)

}

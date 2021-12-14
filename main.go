package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/riceChuang/gamerobot/game"
	"github.com/riceChuang/gamerobot/handler"
	"github.com/riceChuang/gamerobot/util/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.LoadConfig()
	logLevel, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Warnf("log level invalid, set log to info level, configs data:%s", "debug")
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	r := gin.Default()
	r.Static("/img", "./static/img")
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.LoadHTMLGlob("templates/**/*")
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	handler.Router(r)

	//初始化遊戲邏輯
	game.InitGameInstance()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// service connections
		log.Info("Started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// waiting signal and graceful shut down（5 sec over time)
	quit := make(chan os.Signal)
	// syscall SIGINT:ctrl-c, SIGTSTP:ctrl-z, SIGQUIT:ctrl-\
	signal.Notify(quit, os.Interrupt, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTSTP)
	<-quit
	log.Info("Shutdown  ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

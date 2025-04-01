package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sven/sven/contrib/engine"
	"testing"
	"time"
)

func TestNewGinServer(t *testing.T) {
	e := engine.NewEngine("debug", true)
	e.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})
	srv := NewGinServer(e, ":8090")
	go func() {
		time.Sleep(10 * time.Second)
		_ = srv.Stop()
	}()

	err := srv.Start()
	if err != nil {
		panic(err)
	}

}

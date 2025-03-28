package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func TestNewGinServer(t *testing.T) {
	engine := gin.Default()
	engine.Handle("GET", "/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	srv := NewGinServer(engine, ":8090")
	go func() {
		time.Sleep(10 * time.Second)
		_ = srv.Stop()
	}()

	err := srv.Start()
	if err != nil {
		panic(err)
	}

}

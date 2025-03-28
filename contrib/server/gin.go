package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-sven/sven/server"
	"net/http"
	"time"
)

// GinServer 必须实现 server接口
var _ server.Server = (*GinServer)(nil)

//GinServer 实现server 的接口
//去实现 server.Start() server.Stop()

type GinServer struct {
	httpServer *http.Server
}

func (s *GinServer) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *GinServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}

// 使用gin.Engine

func NewGinServer(engine *gin.Engine, addr string) *GinServer {
	return &GinServer{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: engine,
		}}
}

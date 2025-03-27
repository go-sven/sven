package sven

import (
	"errors"
	"github.com/go-sven/sven/log"
	"github.com/go-sven/sven/server"
	"os"
	"os/signal"
)

//想封装个web的骨架
//不知道干啥 上来先定义一个应用

type Application struct {
	id      string
	name    string
	version string
	//应用有了三要素了，应该要有个服务能开始 能结束
	//server Server
	server server.Server
	//接受信号，优雅的停止服务
	signals []os.Signal
}

//抽象出一个能开始 能结束的接口

//type Server interface {
//	Start() error
//	Stop() error
//}//放在server文件夹里

// NewApp new app with opts
func NewApp(opts ...Option) *Application {
	app := new(Application)
	for _, opt := range opts {
		opt(app)
	}
	return app
}

// 初始化应用 加上个option 模式

type Option func(*Application)

// WithId app add id
func WithId(id string) Option {
	return func(app *Application) {
		app.id = id
	}
}

// WithName app add name
func WithName(name string) Option {
	return func(app *Application) {
		app.name = name
	}
}

// WithVersion app add version
func WithVersion(version string) Option {
	return func(app *Application) {
		app.version = version
	}
}

// WithServer app add server
func WithServer(server server.Server) Option {
	return func(app *Application) {
		app.server = server
	}
}

// WithSignals app add signal
func WithSignals(signals []os.Signal) Option {
	return func(app *Application) {
		app.signals = signals
	}
}
func (app *Application) Run() error {
	if app.server == nil {
		return errors.New("no server to run")
	}
	go func(srv server.Server) {
		if err := srv.Start(); err != nil {
			log.Errorf("failed to start server, err: %s", err)
		}
	}(app.server)

	quit := make(chan os.Signal, 8)
	signal.Notify(quit, app.signals...)
	<-quit
	return nil
}

func (app *Application) Stop() error {
	go func(srv server.Server) {
		if err := srv.Stop(); err != nil {
			log.Errorf("failed to stop server, err: %s", err)
		}
	}(app.server)
	return nil
}

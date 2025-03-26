package viper

import (
	"testing"
)

type AppConfig struct {
	Server   *Server   `json:"server" yaml:"server"`
	Database *Database `json:"database" yaml:"database"`
	Cache    *Cache    `json:"cache" yaml:"cache"`
}
type Server struct {
	Addr string `json:"addr" yaml:"addr"`
	Mode string `json:"mode" yaml:"mode"`
}
type Database struct {
	Driver     string `json:"driver"  yaml:"driver"`
	Connection string `json:"connection"  yaml:"connection"`
}

// Cache cache
type Cache struct {
	Addr     string `json:"addr"  yaml:"addr"`
	Password string `json:"password"  yaml:"password"`
	Db       int    `json:"db"  yaml:"db" `
}

func TestViperParse(t *testing.T) {
	t.Run("use viper parse config file", func(t *testing.T) {
		filePath := "../testdata/config.yaml"
		svenViper, err := NewSvenViper(filePath)
		if err != nil {
			t.Error("viper err:", err)
		}
		conf := new(AppConfig)
		err = svenViper.Parse(conf)
		if err != nil {
			t.Error("parse err:", err)
		}
		t.Log(conf.Server.Addr)
		t.Log(conf.Database.Driver)
		t.Log(conf.Cache.Password)
	})

}

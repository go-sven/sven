package viper

import (
	"fmt"
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
	filePath := "../test/config.yaml"
	vc, err := NewViperConfig(filePath)
	if err != nil {
		fmt.Println("viper err:", err)
	}
	conf := new(AppConfig)
	err = vc.Parse(conf)
	if err != nil {
		fmt.Println("parse err:", err)
	}
	fmt.Println(conf.Server.Addr)
	fmt.Println(conf.Database.Driver)
	fmt.Println(conf.Cache.Password)
}

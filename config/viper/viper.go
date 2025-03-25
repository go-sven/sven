package viper

import (
	"fmt"
	"github.com/go-sven/sven/config"
	"github.com/spf13/viper"
	"os"
)

type ViperConfig struct {
	parser *viper.Viper
}

func NewViperConfig(filePath string) (conf config.Config, err error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		return
	}
	if !stat.Mode().IsRegular() {
		return nil, fmt.Errorf("%s is not a regular file", filePath)
	}
	parser := viper.New()
	parser.SetConfigFile(filePath)
	err = parser.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &ViperConfig{parser: parser}, nil
}

func (c *ViperConfig) Parse(obj any) error {
	return c.parser.Unmarshal(obj)
}

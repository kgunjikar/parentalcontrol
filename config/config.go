package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sniffer/logger"
)

type Configuration struct {
	IfName     string
	BufferSize int32
	Filter     string
	HTTPOnly   bool
}

var (
	iface       = "wlp4s0mon"
	ifaceNormal = "wlp4s0"
	buffer      = int32(1600)
	filter      = "tcp and port 80"
)
var normal bool = true

var Config Configuration

func ConfigInit(configName string, paths []string) error {
	viper.SetConfigName(configName)

	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Errorf("Error reading config file:", zap.Error(err))
		return err
	}
	if err := viper.Unmarshal(&Config); err != nil {
		logger.Log.Errorf("Error unmarshalling config:%v", err)
		return err
	}
	return nil
}

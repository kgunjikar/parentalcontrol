package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	IfName string
}

func ConfigInit(configName string, paths []string) error {
	viper.ConfigName(configName)

	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	if err := viper.ReadConfig(); err != nil {
		logger.Errorf("Error reading config file:", zap.Error(err))
		return err
	}
	if err := viper.Unmarshal(&Config); err != nil {
		logger.Errorf("Error unmarshalling config:%v", err)
		return err
	}
	return nil

}

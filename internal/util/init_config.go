package util

import (
	"github.com/spf13/viper"
)

func InitConfig(mode, path string) error {
	// 根据模式读取对应的配置信息
	viper.SetConfigName(mode)   // config file name
	viper.AddConfigPath(path)   // config file path
	err := viper.ReadInConfig() // find and read the config file
	return err
}

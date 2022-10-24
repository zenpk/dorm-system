package config

import (
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/gmp"
	"log"
)

func InitGlobalConfig(mode string) error {
	path, err := gmp.GetModPath()
	if err != nil {
		log.Fatalln(err)
	}
	// read config by mode
	viper.SetConfigName(mode)             // config file mode
	viper.AddConfigPath(path + "configs") // config file path
	err = viper.ReadInConfig()            // find and read the config file
	return err
}

func InitConfig(name string) (*viper.Viper, error) {
	path, err := gmp.GetModPath()
	if err != nil {
		log.Fatalln(err)
	}
	config := viper.New()
	config.SetConfigName(name)
	config.AddConfigPath(path + "configs")
	if err := config.ReadInConfig(); err != nil {
		return nil, err
	}
	return config, nil
}

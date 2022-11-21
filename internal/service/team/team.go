package team

import "github.com/spf13/viper"

type Server struct {
	Config *viper.Viper
	UnimplementedTeamServer
}

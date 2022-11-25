package order

import (
	"context"
	"github.com/spf13/viper"
)

type Server struct {
	Config *viper.Viper
	UnimplementedOrderServer
}

func (s Server) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	return nil, nil
}

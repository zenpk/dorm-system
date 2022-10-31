package dorm

import (
	"context"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/cache"
	"github.com/zenpk/dorm-system/internal/service/common"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Server struct {
	Config *viper.Viper
	UnimplementedDormServer
}

func (s *Server) GetAvailableNum(ctx context.Context, req *EmptyRequest) (*MapReply, error) {
	res, err := cache.All.Dorm.GetAvailableNum(ctx)
	if err != nil {
		return nil, err
	}
	reply := &MapReply{
		Resp: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  ep.ErrOK.Msg,
		},
		Available: res,
	}
	return reply, nil
}

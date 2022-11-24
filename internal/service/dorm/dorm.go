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

func (s Server) GetRemainCnt(ctx context.Context, req *EmptyRequest) (*MapReply, error) {
	res, err := cache.All.Dorm.GetRemainNum(ctx)
	if err != nil {
		return nil, err
	}
	reply := &MapReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "remaining bed counts got successfully",
		},
		RemainCnt: res,
	}
	return reply, nil
}

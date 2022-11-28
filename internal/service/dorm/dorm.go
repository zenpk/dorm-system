package dorm

import (
	"context"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/cache"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/service/common"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Server struct {
	Config *viper.Viper
	UnimplementedDormServer
}

func (s Server) GetRemainCnt(ctx context.Context, _ *EmptyRequest) (*MapReply, error) {
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

func (s Server) GetAll(ctx context.Context, _ *EmptyRequest) (*GetAllReply, error) {
	buildings, err := dal.Table.Building.FindAllEnabled(ctx)
	if err != nil {
		return nil, err
	}
	infos := make([]*BuildingInfo, 0)
	for _, b := range buildings {
		info := &BuildingInfo{
			Num:    b.Num,
			Info:   b.Info,
			ImgUrl: b.ImgUrl,
		}
		infos = append(infos, info)
	}
	reply := &GetAllReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "building information got successfully",
		},
		Infos: infos,
	}
	return reply, nil
}

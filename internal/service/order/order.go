package order

import (
	"context"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/service/common"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Server struct {
	Config *viper.Viper
	UnimplementedOrderServer
}

func (s *Server) Submit(ctx context.Context, req *OrderRequest) (*OrderReply, error) {
	ids := make([]uint64, 0)
	ids = append(ids, req.StudentId1)
	if req.StudentId2 != 0 {
		ids = append(ids, req.StudentId2)
	}
	if req.StudentId3 != 0 {
		ids = append(ids, req.StudentId3)
	}
	if req.StudentId4 != 0 {
		ids = append(ids, req.StudentId4)
	}
	// check if already has dorm
	users := make([]*dal.UserInfo, 0)
	userTable := new(dal.UserInfo)
	for _, id := range ids {
		user, err := userTable.FindById(ctx, id)
		if err != nil {
			return nil, err
		}
		if user.DormId != 0 {
			errPack := ep.ErrInputBody
			errPack.Msg = "someone already has a dorm"
			return nil, errPack
		}
		users = append(users, user)
	}
	num := len(users)
	dormTable := new(dal.Dorm)
	dorm, err := dormTable.Allocate(ctx, num)
	if err != nil {
		return nil, err
	}
	for i := range users {
		users[i].DormId = dorm.Id
		if err := userTable.Update(ctx, users[i]); err != nil {
			return nil, err
		}
	}
	reply := &OrderReply{
		Resp: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "submit success",
		},
	}
	return reply, nil
}

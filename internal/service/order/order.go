package order

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/cache"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/service/common"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/ep"
	"github.com/zenpk/dorm-system/pkg/zap"
)

type Server struct {
	Config *viper.Viper
	UnimplementedOrderServer
}

func (s Server) Submit(ctx context.Context, req *SubmitRequest) (*SubmitReply, error) {
	zap.Logger.Infof("handling order %v", req.Code)
	building, err := s.prepareMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	// check if building is enabled
	if building.Enabled == false {
		// not enabled, fail this order
		return nil, s.Fail(ctx, req, "building is not enabled")
	}
	// check if building has remaining beds
	buildingRemain, err := cache.Redis.HGet(ctx, "remain", util.UintToString(building.Id)).Uint64()
	if buildingRemain <= 0 {
		return nil, s.Fail(ctx, req, "this building doesn't have remaining bed anymore")
	}
	order := &dal.Order{
		BuildingId: building.Id,
		DormId:     0,
		TeamId:     req.Team.Id,
		Code:       req.Code,
		Info:       "success",
		Success:    true,
	}
	// get a mutex lock with buildingId in name
	mutex := cache.All.RedSync.NewMutex(s.Config.GetString("mutex_prefix") + util.UintToString(building.Id))
	err = mutex.LockContext(ctx)
	for err != nil { // block the goroutine until get the mutex
		select {
		case <-ctx.Done():
			// timeout
			return nil, errors.New("unable to get mutex before timeout")
		default:
			err = mutex.LockContext(ctx)
		}
	}
	memberCnt := uint64(len(req.Team.Members)) + 1 // number of members in this team
	// allocate a suitable dorm
	dorm, err := dal.Table.Dorm.Allocate(ctx, building.Id, memberCnt, req.Team.Gender)
	if err != nil {
		_, _ = mutex.Unlock()
		return nil, err
	}
	order.DormId = dorm.Id
	// write in the order and decrease dorm's remain count
	//if err := dal.Table.Order.TransCreateAndDecreaseDormRemainCnt(ctx, order, dorm, memberCnt); err != nil {
	//	_, _ = mutex.Unlock()
	//	return nil, err
	//}
	// decrease redis count, buildingRemain needs to be fetched again
	// no need to rollback even if any error occurs
	buildingRemain, err = cache.Redis.HGet(ctx, "remain", util.UintToString(building.Id)).Uint64()
	if err != nil {
		_, _ = mutex.Unlock()
		return nil, err
	}
	buildingRemain -= memberCnt
	//if err := cache.Redis.HSet(ctx, "remain", building.Id, buildingRemain).Err(); err != nil {
	//	_, _ = mutex.Unlock()
	//	return nil, err
	//}
	if _, err := mutex.Unlock(); err != nil {
		return nil, err
	}
	return nil, nil
}

// Fail an order, directly write into database as failed
func (s Server) Fail(ctx context.Context, req *SubmitRequest, info string) error {
	building, err := s.prepareMessage(ctx, req)
	if err != nil {
		return err
	}
	order := &dal.Order{
		BuildingId: building.Id,
		DormId:     0,
		TeamId:     req.Team.Id,
		Code:       req.Code,
		Info:       info,
		Success:    false,
	}
	return dal.Table.Order.Create(ctx, order)
}

func (s Server) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	orders, err := dal.Table.Order.FindAllByTeamIdWithDeleted(ctx, req.TeamId)
	if err != nil {
		return nil, err
	}
	resp := &GetReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "successfully got order information",
		},
		Orders: make([]*OrderInfo, 0),
	}
	for _, order := range orders {
		info, err := s.dalOrderToInfo(ctx, order)
		if err != nil {
			return nil, err
		}
		resp.Orders = append(resp.Orders, info)
	}
	return resp, nil
}

func (s Server) Delete(ctx context.Context, req *DeleteRequest) (*DeleteReply, error) {
	order, err := dal.Table.Order.FindByIdWithDeleted(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	dorm, err := dal.Table.Dorm.FindById(ctx, order.DormId)
	if err != nil {
		return nil, err
	}
	// get a mutex lock with buildingId in name
	mutex := cache.All.RedSync.NewMutex(s.Config.GetString("mutex_prefix") + util.UintToString(order.BuildingId))
	err = mutex.LockContext(ctx)
	for err != nil { // block the goroutine until get the mutex
		select {
		case <-ctx.Done():
			// timeout
			return nil, errors.New("unable to get mutex before timeout")
		default:
			err = mutex.LockContext(ctx)
		}
	}
	memberCnt := uint64(len(req.Team.Members)) + 1 // number of members in this team
	// delete the order
	if err := dal.Table.Order.TransDeleteAndIncreaseDormRemainCnt(ctx, order, dorm, memberCnt); err != nil {
		_, _ = mutex.Unlock()
		return nil, err
	}
	// increase redis count
	buildingRemain, err := cache.Redis.HGet(ctx, "remain", util.UintToString(order.BuildingId)).Uint64()
	if err != nil {
		_, _ = mutex.Unlock()
		return nil, err
	}
	buildingRemain += memberCnt
	if err := cache.Redis.HSet(ctx, "remain", order.BuildingId, buildingRemain).Err(); err != nil {
		_, _ = mutex.Unlock()
		return nil, err
	}
	if _, err := mutex.Unlock(); err != nil {
		return nil, err
	}
	reply := &DeleteReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "successfully deleted the order",
		},
	}
	return reply, nil
}

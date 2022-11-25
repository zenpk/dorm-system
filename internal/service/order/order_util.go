package order

import (
	"context"
	"errors"
	"github.com/zenpk/dorm-system/internal/dal"
	"gorm.io/gorm"
)

// prepareMessage from Kafka, return the target building
func (s Server) prepareMessage(ctx context.Context, req *SubmitRequest) (building *dal.Building, err error) {
	// check if order already handled
	_, err = dal.Table.Order.FindByCodeWithDeleted(ctx, req.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // errors other than RecordNotFound
		return nil, err
	}
	if err == nil { // already have this code
		return nil, errors.New("cannot create multiple orders with the same code")
	}
	building, err = dal.Table.Building.FindByNum(ctx, req.BuildingNum)
	if err != nil {
		return nil, err
	}
	return building, nil
}

// dalOrderToInfo convert *dal.Order to *OrderInfo
func (s Server) dalOrderToInfo(ctx context.Context, order *dal.Order) (info *OrderInfo, err error) {
	building, err := dal.Table.Building.FindById(ctx, order.BuildingId)
	if err != nil {
		return nil, err
	}
	dorm, err := dal.Table.Dorm.FindById(ctx, order.DormId)
	if err != nil {
		return nil, err
	}
	info = &OrderInfo{
		Id:          order.Id,
		BuildingNum: building.Num,
		DormNum:     dorm.Num,
		Info:        order.Info,
		Success:     order.Success,
		Deleted:     order.Deleted.Valid,
	}
	return
}

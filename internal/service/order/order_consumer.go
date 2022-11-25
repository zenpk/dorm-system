package order

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
	"github.com/zenpk/dorm-system/internal/dal"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

func Submit(ctx context.Context, message *sarama.ConsumerMessage) error {
	var req SubmitRequest
	if err := proto.Unmarshal(message.Value, &req); err != nil {
		return err
	}
	// check if order already handled
	_, err := dal.Table.Order.FindByCode(ctx, req.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // errors other than RecordNotFound
		return err
	}
	if err == nil { // already have this code
		return errors.New("cannot create multiple orders with the same code")
	}
	building, err := dal.Table.Building.FindByNum(ctx, req.BuildingNum)
	if err != nil {
		return err
	}
	order := &dal.Order{
		BuildingId: building.Id,
		DormId:     0,
		TeamId:     req.Team.Id,
		Code:       req.Code,
		Success:    false,
	}
	return dal.Table.Order.Create(ctx, order)
}

// Fail an order, directly write into database as failed
func Fail(ctx context.Context, message *sarama.ConsumerMessage) error {
	var req SubmitRequest
	if err := proto.Unmarshal(message.Value, &req); err != nil {
		return err
	}
	// check if order already handled
	_, err := dal.Table.Order.FindByCode(ctx, req.Code)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // errors other than RecordNotFound
		return err
	}
	if err == nil { // already have this code
		return errors.New("cannot create multiple orders with the same code")
	}
	building, err := dal.Table.Building.FindByNum(ctx, req.BuildingNum)
	if err != nil {
		return err
	}
	order := &dal.Order{
		BuildingId: building.Id,
		DormId:     0,
		TeamId:     req.Team.Id,
		Code:       req.Code,
		Success:    false,
	}
	return dal.Table.Order.Create(ctx, order)
}

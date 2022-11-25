package dal

import (
	"context"
	"gorm.io/gorm"
)

type Order struct {
	Id         uint64 `gorm:"primaryKey"`
	BuildingId uint64 `gorm:"not null; index"`
	DormId     uint64 `gorm:"not null; default:0; index"`
	TeamId     uint64 `gorm:"not null; index"`
	Code       string `gorm:"not null; unique; index"`
	Info       string
	Success    bool           `gorm:"not null; default:0"`
	Deleted    gorm.DeletedAt `gorm:"index"`
}

func (o Order) FindById(ctx context.Context, id uint64) (order *Order, err error) {
	return order, DB.WithContext(ctx).Take(&order, id).Error
}

// FindByCode is to make sure code is unique, this will include deleted records
func (o Order) FindByCode(ctx context.Context, code string) (order *Order, err error) {
	return order, DB.WithContext(ctx).Unscoped().Where("code = ?", code).Take(&order).Error
}

func (o Order) FindAllByTeamId(ctx context.Context, id uint64) (orders []*Order, err error) {
	return orders, DB.WithContext(ctx).Where("team_id = ?", id).Find(&orders).Error
}

func (o Order) FindSuccessByTeamId(ctx context.Context, id uint64) (order *Order, err error) {
	return order, DB.WithContext(ctx).Where("success = true AND team_id = ?", id).Take(&order).Error
}

func (o Order) Create(ctx context.Context, order *Order) error {
	return DB.WithContext(ctx).Create(&order).Error
}

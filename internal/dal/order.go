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
	Code       string `gorm:"not null; index"`
	Info       string
	Success    bool           `gorm:"not null; default:0"`
	Deleted    gorm.DeletedAt `gorm:"index"`
}

func (o Order) FindById(ctx context.Context, id uint64) (order *Order, err error) {
	return order, DB.WithContext(ctx).Take(&order, id).Error
}

func (o Order) FindSuccessByTeamId(ctx context.Context, id uint64) (order *Order, err error) {
	return order, DB.WithContext(ctx).Where("success = true AND id = ?", id).Take(&order).Error
}

func (o Order) Create(ctx context.Context, order *Order) error {
	return DB.WithContext(ctx).Create(&order).Error
}

package dal

import (
	"context"
	"gorm.io/gorm"
)

type Order struct {
	Id         uint64 `gorm:"primaryKey"`
	BuildingId uint64 `gorm:"not null"`
	DormId     uint64 `gorm:"not null; default:0"`
	TeamId     uint64 `gorm:"not null"`
	Code       string `gorm:"not null; unique"`
	Info       string
	Success    bool `gorm:"not null; default:0"`
	Deleted    gorm.DeletedAt
}

func (o Order) FindById(ctx context.Context, id uint64) (order *Order, err error) {
	return order, DB.WithContext(ctx).Take(&order, id).Error
}

func (o Order) FindByIdWithDeleted(ctx context.Context, id uint64) (order *Order, err error) {
	return order, DB.WithContext(ctx).Unscoped().Take(&order, id).Error
}

// FindByCodeWithDeleted is to make sure code is unique, this will include deleted records
func (o Order) FindByCodeWithDeleted(ctx context.Context, code string) (order *Order, err error) {
	return order, DB.WithContext(ctx).Unscoped().Where("code = ?", code).Take(&order).Error
}

func (o Order) FindAllByTeamIdWithDeleted(ctx context.Context, id uint64) (orders []*Order, err error) {
	return orders, DB.WithContext(ctx).Unscoped().Where("team_id = ?", id).Find(&orders).Error
}

func (o Order) FindSuccessByTeamId(ctx context.Context, id uint64) (order *Order, err error) {
	return order, DB.WithContext(ctx).Where("success = true AND team_id = ?", id).Take(&order).Error
}

// TransCreateAndDecreaseDormRemainCnt create an order and decrease the corresponding dorm's remain count
func (o Order) TransCreateAndDecreaseDormRemainCnt(ctx context.Context, order *Order, dorm *Dorm, memberCnt uint64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		dorm.RemainCnt -= memberCnt
		if err := dorm.Update(ctx, dorm); err != nil {
			return err
		}
		if err := o.Create(ctx, order); err != nil {
			return err
		}
		return nil
	})
}

// TransDeleteAndIncreaseDormRemainCnt delete an order and increase the corresponding dorm's remain count
func (o Order) TransDeleteAndIncreaseDormRemainCnt(ctx context.Context, order *Order, dorm *Dorm, memberCnt uint64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		dorm.RemainCnt += memberCnt
		if err := Table.Dorm.Update(ctx, dorm); err != nil {
			return err
		}
		order.Info += " (deleted)"
		if err := DB.WithContext(ctx).Unscoped().Save(&order).Error; err != nil {
			return err
		}
		if err := o.Delete(ctx, order); err != nil {
			return err
		}
		return nil
	})
}

func (o Order) Update(ctx context.Context, order *Order) error {
	return DB.WithContext(ctx).Save(&order).Error
}

func (o Order) Create(ctx context.Context, order *Order) error {
	return DB.WithContext(ctx).Create(&order).Error
}

func (o Order) Delete(ctx context.Context, order *Order) error {
	return DB.WithContext(ctx).Delete(&order).Error
}

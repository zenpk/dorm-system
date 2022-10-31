package dal

import (
	"context"
	"errors"
)

type Dorm struct {
	Id         uint64 `gorm:"primaryKey" json:"-"`
	DormId     uint64 `gorm:"unique; not null; index" json:"dormId,omitempty"`
	BuildingId uint64 `gorm:"not null; index" json:"buildingId,omitempty"`
	Gender     string `gorm:"not null" json:"gender,omitempty"`
	Available  uint64 `gorm:"not null" json:"available,omitempty"`
	BedNum     uint64 `gorm:"not null" json:"bedNum,omitempty"`
	Info       string `json:"info,omitempty"`
}

func (d *Dorm) SumAvailableByBuildingId(ctx context.Context, id uint64) (int64, error) {
	building, err := Table.Building.FindById(ctx, id)
	if err != nil {
		return 0, err
	}
	if building.IsAvailable == false {
		return 0, errors.New("building unavailable")
	}
	var sum int64
	return sum, DB.WithContext(ctx).Model(&Dorm{}).Where("building_id = ?", id).Select("SUM(available)").Row().Scan(&sum)
}

func (d *Dorm) Allocate(ctx context.Context, buildingId uint64, num uint64, gender string) (*Dorm, error) {
	dorm := new(Dorm)
	// TODO: unable buildings
	err := DB.WithContext(ctx).Where("building_id = ? AND available > ? AND gender = ?", buildingId, num, gender).First(&dorm).Error
	if err != nil {
		return nil, err
	}
	dorm.Available -= num
	return dorm, d.Update(ctx, dorm)
}

func (d *Dorm) Update(ctx context.Context, dorm *Dorm) error {
	return DB.WithContext(ctx).Save(&dorm).Error
}

//func (d *Dorm) SumBedNumByBuildingId(ctx context.Context, id uint32) (int64, error) {
//	var sum int64
//	return sum, DB.WithContext(ctx).Model(&Dorm{}).Select("SUM(bed_num)").Row().Scan(&sum)
//}

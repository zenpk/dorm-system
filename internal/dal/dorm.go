package dal

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Dorm struct {
	Id         uint64 `gorm:"primaryKey" json:"-"`
	Num        string `gorm:"size:10; not null; unique" json:"num,omitempty"`
	BuildingId uint64 `gorm:"not null" json:"buildingId,omitempty"`
	Gender     string `gorm:"size:10; not null" json:"gender,omitempty"`
	RemainCnt  uint64 `gorm:"not null" json:"remainCnt,omitempty"`
	BedCnt     uint64 `gorm:"not null" json:"bedCnt,omitempty"`
	Info       string `json:"info,omitempty"`
	Enabled    bool   `gorm:"not null" json:"enabled,omitempty"`
	Deleted    gorm.DeletedAt
}

func (d Dorm) FindById(ctx context.Context, id uint64) (dorm *Dorm, err error) {
	return dorm, DB.WithContext(ctx).Take(&dorm, id).Error
}

func (d Dorm) FindAllEnabled(ctx context.Context) (dorms []*Dorm, err error) {
	return dorms, DB.WithContext(ctx).Where("enabled = true").Find(&dorms).Error
}

func (d Dorm) SumRemainCntByBuildingId(ctx context.Context, id uint64) (sum int64, err error) {
	building, err := Table.Building.FindById(ctx, id)
	if err != nil {
		return 0, err
	}
	if building.Enabled == false {
		return 0, errors.New("building is not enabled")
	}
	return sum, DB.WithContext(ctx).Model(&Dorm{}).Select("SUM(remain_cnt)").Where("building_id = ?", id).Scan(&sum).Error
}

// Allocate find a suitable dorm for the order
func (d Dorm) Allocate(ctx context.Context, buildingId uint64, memberCnt uint64, gender string) (dorm *Dorm, err error) {
	return dorm, DB.WithContext(ctx).Where("building_id = ? AND remain_cnt > ? AND gender = ? AND enabled = true", buildingId, memberCnt, gender).Take(&dorm).Error
}

func (d Dorm) Update(ctx context.Context, dorm *Dorm) error {
	return DB.WithContext(ctx).Save(&dorm).Error
}

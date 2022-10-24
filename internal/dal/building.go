package dal

import "context"

type Building struct {
	Id          uint64 `gorm:"primaryKey" json:"-"`
	BuildingId  uint64 `gorm:"unique; not null; index" json:"buildingId,omitempty"`
	IsAvailable bool   `gorm:"not null" json:"isAvailable,omitempty"`
	Info        string `json:"info,omitempty"`
}

func (b *Building) FindById(ctx context.Context, id uint64) (*Building, error) {
	building := new(Building)
	return building, DB.WithContext(ctx).First(&building, id).Error
}

func (b *Building) FindAllAvailable(ctx context.Context) ([]*Building, error) {
	var buildings []*Building
	return buildings, DB.WithContext(ctx).Where("is_available = true").Find(&buildings).Error
}

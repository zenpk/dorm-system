package dal

import "context"

type Building struct {
	Id      uint64 `gorm:"primaryKey" json:"-"`
	Num     string `gorm:"size:10; not null; unique; index" json:"num,omitempty"`
	Enabled bool   `gorm:"not null; default:1" json:"enabled,omitempty"`
	Info    string `json:"info,omitempty"`
}

func (b *Building) FindById(ctx context.Context, id uint64) (*Building, error) {
	building := new(Building)
	return building, DB.WithContext(ctx).First(&building, id).Error
}

func (b *Building) FindAllAvailable(ctx context.Context) ([]*Building, error) {
	var buildings []*Building
	return buildings, DB.WithContext(ctx).Where("is_available = true").Find(&buildings).Error
}

func (b *Building) FindAllAvailableIds(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	return ids, DB.WithContext(ctx).Model(&Building{}).Where("is_available = true").Distinct().Pluck("building_id", &ids).Error
}

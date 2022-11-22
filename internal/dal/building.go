package dal

import (
	"context"
	"gorm.io/gorm"
)

type Building struct {
	Id       uint64         `gorm:"primaryKey" json:"-"`
	Num      string         `gorm:"size:10; not null; unique; index" json:"num,omitempty"`
	ImageUrl string         `json:"imageUrl,omitempty"`
	Info     string         `json:"info,omitempty"`
	Enabled  bool           `gorm:"not null; default:1; index" json:"enabled,omitempty"`
	Deleted  gorm.DeletedAt `gorm:"index"`
}

func (b *Building) FindById(ctx context.Context, id uint64) (*Building, error) {
	building := new(Building)
	return building, DB.WithContext(ctx).First(&building, id).Error
}

func (b *Building) FindAllAvailable(ctx context.Context) ([]*Building, error) {
	var buildings []*Building
	return buildings, DB.WithContext(ctx).Where("enabled = true").Find(&buildings).Error
}

func (b *Building) PluckAllAvailableIds(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	return ids, DB.WithContext(ctx).Model(&Building{}).Where("enabled = true").Distinct().Pluck("id", &ids).Error
}

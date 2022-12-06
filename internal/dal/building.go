package dal

import (
	"context"
	"gorm.io/gorm"
)

type Building struct {
	Id      uint64         `gorm:"primaryKey" json:"-"`
	Num     string         `gorm:"size:10; not null; unique; index" json:"num,omitempty"`
	Info    string         `json:"info,omitempty"`
	ImgUrl  string         `json:"imgUrl,omitempty"`
	Enabled bool           `gorm:"not null; index" json:"enabled,omitempty"`
	Deleted gorm.DeletedAt `gorm:"index"`
}

func (b Building) FindById(ctx context.Context, id uint64) (building *Building, err error) {
	return building, DB.WithContext(ctx).Take(&building, id).Error
}

func (b Building) FindByNum(ctx context.Context, num string) (building *Building, err error) {
	return building, DB.WithContext(ctx).Where("num = ?", num).Take(&building).Error
}

func (b Building) FindAllEnabled(ctx context.Context) (buildings []*Building, err error) {
	return buildings, DB.WithContext(ctx).Where("enabled = true").Find(&buildings).Error
}

func (b Building) PluckAllEnabledIds(ctx context.Context) (ids []uint64, err error) {
	return ids, DB.WithContext(ctx).Model(&Building{}).Select("id").Where("enabled = true").Scan(&ids).Error
}

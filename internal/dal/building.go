package dal

import "context"

type Building struct {
	Id      uint64 `gorm:"primaryKey" json:"-"`
	Num     string `gorm:"size:10; not null; unique; index" json:"num,omitempty"`
	Enabled bool   `gorm:"not null; default:1" json:"enabled,omitempty"`
	Info    string `json:"info,omitempty"`
}

func (b *Building) FindById(ctx context.Context, id uint64) (building *Building, err error) {
	err = DB.WithContext(ctx).First(&building, id).Error
	return
}

func (b *Building) FindAllAvailable(ctx context.Context) (buildings []*Building, err error) {
	err = DB.WithContext(ctx).Where("enabled = true").Find(&buildings).Error
	return
}

func (b *Building) FindAllAvailableIds(ctx context.Context) (ids []uint64, err error) {
	err = DB.WithContext(ctx).Model(&Building{}).Where("enabled = true").Distinct().Pluck("id", &ids).Error
	return
}

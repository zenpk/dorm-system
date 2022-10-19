package dal

type Building struct {
	Id         uint32 `gorm:"primaryKey"`
	BuildingId uint32 `gorm:"unique; not null; index"`
	// use sum instead
	//Available  uint32 `gorm:"not null"`
	//AllNum     uint32 `gorm:"not null"`
	IsAvailable bool `gorm:"not null"`
	Info        string
}

func (b *Building) FindById(id uint32) (*Building, error) {
	building := new(Building)
	return building, DB.First(&building, id).Error
}

func (b *Building) FindAllAvailable() ([]*Building, error) {
	var buildings []*Building
	return buildings, DB.Where("is_available = true").Find(&buildings).Error
}

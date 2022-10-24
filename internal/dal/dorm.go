package dal

import "errors"

type Dorm struct {
	Id         uint32 `gorm:"primaryKey"`
	DormId     uint32 `gorm:"unique; not null; index"`
	BuildingId uint32 `gorm:"not null; index"`
	Gender     string `gorm:"not null"`
	Available  uint32 `gorm:"not null"`
	BedNum     uint32 `gorm:"not null"`
	Info       string
}

func (d *Dorm) SumAvailableByBuildingId(id uint32) (int64, error) {
	table := new(Building)
	building, err := table.FindById(id)
	if err != nil {
		return 0, err
	}
	if building.IsAvailable == false {
		return 0, errors.New("building unavailable")
	}
	var sum int64
	return sum, DB.Model(&Dorm{}).Select("SUM(available)").Row().Scan(&sum)
}

//func (d *Dorm) SumBedNumByBuildingId(id uint32) (int64, error) {
//	var sum int64
//	return sum, DB.Model(&Dorm{}).Select("SUM(bed_num)").Row().Scan(&sum)
//}

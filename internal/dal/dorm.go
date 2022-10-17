package dal

type dorm struct {
	Id      int64 `gorm:"primaryKey"`
	DormNum int64 `gorm:""`
}

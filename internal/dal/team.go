package dal

type Team struct {
	Id      uint64 `gorm:"primaryKey"`
	Code    string `gorm:"not null; unique; index"`
	Gender  string `gorm:"size:10; not null"`
	OwnerId uint64 `gorm:"not null; unique; index"`
	Deleted bool   `gorm:"not null; default:0; index"`
}

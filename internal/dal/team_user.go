package dal

type TeamUser struct {
	Id      uint64 `gorm:"primaryKey"`
	TeamId  uint64 `gorm:"not null; index"`
	UserId  uint64 `gorm:"not null; index"`
	Deleted bool   `gorm:"not null; default:0; index"`
}

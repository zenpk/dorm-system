package dal

import "gorm.io/gorm"

type TeamUser struct {
	Id      uint64         `gorm:"primaryKey"`
	TeamId  uint64         `gorm:"not null; index"`
	UserId  uint64         `gorm:"not null; index"`
	Deleted gorm.DeletedAt `gorm:"index"`
}

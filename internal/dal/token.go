package dal

import "time"

type Token struct {
	Id           uint64     `gorm:"primaryKey" json:"-"`
	RefreshToken string     `gorm:"not null" json:"refreshToken,omitempty"`
	Uid          uint64     `gorm:"not null" json:"-"`
	CreateTime   *time.Time `gorm:"not null" json:"createTime,omitempty"`
	ExpTime      *time.Time `gorm:"not null" json:"expTime,omitempty"`
	Deleted      bool       `gorm:"not null; default:0" json:"deleted,omitempty"`
}

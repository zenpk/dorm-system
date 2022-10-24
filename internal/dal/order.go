package dal

import "context"

type Order struct {
	Id         uint64 `gorm:"primaryKey" json:"-"`
	DormId     uint64 `gorm:"not null" json:"dormId"`
	StudentId1 uint64 `gorm:"unique; not null" json:"studentId1,omitempty"`
	StudentId2 uint64 `gorm:"unique" json:"studentId2,omitempty"`
	StudentId3 uint64 `gorm:"unique" json:"studentId3,omitempty"`
	StudentId4 uint64 `gorm:"unique" json:"studentId4,omitempty"`
}

func (o *Order) Create(ctx context.Context, order *Order) error {
	return DB.WithContext(ctx).Create(&order).Error
}

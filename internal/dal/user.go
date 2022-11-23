package dal

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	Id         uint64         `gorm:"primaryKey" json:"-"`
	StudentNum string         `gorm:"not null; unique; index" json:"studentNum,omitempty"`
	Name       string         `gorm:"size:20; not null; index" json:"name,omitempty"`
	Gender     string         `gorm:"size:10; not null; index" json:"gender,omitempty"`
	Role       int32          `gorm:"not null; default:1; index" json:"-"`
	Deleted    gorm.DeletedAt `gorm:"index"`
}

func (u *User) FindById(ctx context.Context, userId uint64) (user *User, err error) {
	return user, DB.WithContext(ctx).Take(&user, userId).Error
}

func (u *User) FindAll(ctx context.Context) (users []*User, err error) {
	return users, DB.WithContext(ctx).Find(&users).Error
}

func (u *User) FindByStudentNum(ctx context.Context, studentNum string) (user *User, err error) {
	return user, DB.WithContext(ctx).Where("student_num = ?", studentNum).Take(&user).Error
}

func (u *User) FindAllByIds(ctx context.Context, ids []uint64) (users []*User, err error) {
	return users, DB.WithContext(ctx).Where("id IN ?", ids).Find(&users).Error
}

func (u *User) Create(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(&user).Error
}

func (u *User) Update(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Save(&user).Error
}

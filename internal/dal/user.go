package dal

import (
	"context"
)

type User struct {
	Id         uint64 `gorm:"primaryKey" json:"-"`
	StudentNum string `gorm:"not null; unique; index" json:"studentNum,omitempty"`
	Name       string `gorm:"size:20; not null; index" json:"name,omitempty"`
	Gender     string `gorm:"size:10; not null; index" json:"gender,omitempty"`
	Role       int32  `gorm:"not null; default:0; index" json:"-"`
	Deleted    uint64 `gorm:"not null; default:0; index"`
}

func (u *User) FindById(ctx context.Context, userId uint64) (*User, error) {
	user := new(User)
	return user, DB.WithContext(ctx).First(&user, userId).Error
}

func (u *User) FindAll(ctx context.Context) ([]*User, error) {
	var users []*User
	return users, DB.WithContext(ctx).Find(&users).Error
}

func (u *User) FindByCredentialId(ctx context.Context, userId uint64) (*User, error) {
	user := new(User)
	return user, DB.WithContext(ctx).Where("user_id = ?", userId).First(&user).Error
}

func (u *User) FindByStudentNum(ctx context.Context, studentNum string) (*User, error) {
	user := new(User)
	return user, DB.WithContext(ctx).Where("student_num = ?", studentNum).First(&user).Error
}

func (u *User) Create(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(&user).Error
}

func (u *User) Update(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Save(&user).Error
}

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
	userInfo := new(User)
	return userInfo, DB.WithContext(ctx).First(&userInfo, userId).Error
}

func (u *User) FindAll(ctx context.Context) ([]*User, error) {
	var userInfos []*User
	return userInfos, DB.WithContext(ctx).Find(&userInfos).Error
}

func (u *User) FindByCredentialId(ctx context.Context, userId uint64) (*User, error) {
	userInfo := new(User)
	return userInfo, DB.WithContext(ctx).Where("user_id = ?", userId).First(&userInfo).Error
}

func (u *User) FindByStudentNum(ctx context.Context, studentNum string) (*User, error) {
	userInfo := new(User)
	return userInfo, DB.WithContext(ctx).Where("student_num = ?", studentNum).First(&userInfo).Error
}

func (u *User) Create(ctx context.Context, info *User) error {
	return DB.WithContext(ctx).Create(&info).Error
}

func (u *User) Update(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Save(&user).Error
}

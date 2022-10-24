package dal

import (
	"context"
)

type UserInfo struct {
	Id               uint64 `gorm:"primaryKey" json:"-"`
	UserCredentialId uint64 `gorm:"unique; not null; index" json:"-"`
	Username         string `gorm:"unique; not null; index" json:"username,omitempty"`
	StudentId        uint64 `gorm:"unique; not null; index" json:"studentId,omitempty"`
	Gender           string `gorm:"not null" json:"gender,omitempty"`
	Name             string `gorm:"not null; index;" json:"name,omitempty"`
	DormId           uint64 `gorm:"not null; default:0" json:"dormId,omitempty"`
}

func (u *UserInfo) FindById(ctx context.Context, userId uint64) (*UserInfo, error) {
	userInfo := new(UserInfo)
	return userInfo, DB.WithContext(ctx).First(&userInfo, userId).Error
}

func (u *UserInfo) FindAll(ctx context.Context) ([]*UserInfo, error) {
	var userInfos []*UserInfo
	return userInfos, DB.WithContext(ctx).Find(&userInfos).Error
}

func (u *UserInfo) FindByUserCredentialId(ctx context.Context, userId uint64) (*UserInfo, error) {
	userInfo := new(UserInfo)
	return userInfo, DB.WithContext(ctx).Where("user_id = ?", userId).First(&userInfo).Error
}

func (u *UserInfo) FindByStudentId(ctx context.Context, studentId uint64) (*UserInfo, error) {
	userInfo := new(UserInfo)
	return userInfo, DB.WithContext(ctx).Where("student_id = ?", studentId).First(&userInfo).Error
}

func (u *UserInfo) Create(ctx context.Context, info *UserInfo) error {
	return DB.WithContext(ctx).Create(&info).Error
}

func (u *UserInfo) Update(ctx context.Context, user *UserInfo) error {
	return DB.WithContext(ctx).Save(&user).Error
}

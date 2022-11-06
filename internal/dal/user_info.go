package dal

import (
	"context"
)

type UserInfo struct {
	Id           uint64 `gorm:"primaryKey" json:"-"`
	CredentialId uint64 `gorm:"not null; unique; index" json:"-"`
	Username     string `gorm:"not null; unique; index" json:"username,omitempty"`
	StudentNum   string `gorm:"not null; unique; index" json:"studentNum,omitempty"`
	Name         string `gorm:"size:20; not null; index" json:"name,omitempty"`
	Gender       string `gorm:"size:10; not null; index" json:"gender,omitempty"`
	Role         int32  `gorm:"not null; default:0; index" json:"-"`
	DormId       uint64 `gorm:"not null; default:0; index" json:"dormId,omitempty"`
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

func (u *UserInfo) FindByStudentId(ctx context.Context, studentId string) (*UserInfo, error) {
	userInfo := new(UserInfo)
	return userInfo, DB.WithContext(ctx).Where("student_id = ?", studentId).First(&userInfo).Error
}

func (u *UserInfo) Create(ctx context.Context, info *UserInfo) error {
	return DB.WithContext(ctx).Create(&info).Error
}

func (u *UserInfo) Update(ctx context.Context, user *UserInfo) error {
	return DB.WithContext(ctx).Save(&user).Error
}

package dal

import "github.com/google/uuid"

type UserInfo struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	UserId    int64  `gorm:"unique; not null; index" json:"-"`
	Username  string `gorm:"unique; not null; index" json:"username"`
	StudentId int64  `gorm:"unique; not null; index" json:"studentId"`
	Name      string `gorm:"not null; index;" json:"name"`
	UUID      string `gorm:"unique; not null;" json:"uuid"`
}

func (u *UserInfo) FindById(userId int64) (UserInfo, error) {
	var userInfo UserInfo
	return userInfo, DB.First(&userInfo, userId).Error
}

func (u *UserInfo) FindAll() ([]UserInfo, error) {
	var userInfos []UserInfo
	return userInfos, DB.Find(&userInfos).Error
}

func (u *UserInfo) FindByUserId(userId int64) (UserInfo, error) {
	var userInfo UserInfo
	return userInfo, DB.Where("user_id = ?", userId).First(&userInfo).Error
}

func (u *UserInfo) FindByStudentId(studentId int64) (UserInfo, error) {
	var userInfo UserInfo
	return userInfo, DB.Where("student_id = ?", studentId).First(&userInfo).Error
}

// Create a new record with randomly generated UUID
func (u *UserInfo) Create(info *UserInfo) error {
	info.UUID = uuid.New().String()
	return DB.Create(info).Error
}

package dal

// UserInfo 用户详细信息表
type UserInfo struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	UserId      int64  `gorm:"unique; index; not null;"`
	Username    string `gorm:"index; not null;" json:"username"`
	Email       string `gorm:"unique; index; not null;" json:"email"`
	IsAuthority int64  `gorm:"not null" json:"isauthority"`
}

// FindById 根据 id 查找用户信息
func (u *UserInfo) FindById(userId int64) (UserInfo, error) {
	var userInfo UserInfo
	return userInfo, DB.First(&userInfo, userId).Error
}

// FindAll 获取全部用户信息
func (u *UserInfo) FindAll() ([]UserInfo, error) {
	var userInfos []UserInfo
	return userInfos, DB.Find(&userInfos).Error
}

// FindByUserId 根据 userId 查找用户信息
func (u *UserInfo) FindByUserId(userId int64) (UserInfo, error) {
	var userInfo UserInfo
	return userInfo, DB.Where("user_id = ?", userId).First(&userInfo).Error
}

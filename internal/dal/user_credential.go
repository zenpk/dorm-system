package dal

// UserCredential 用户登录信息表
type UserCredential struct {
	Id       int64  `gorm:"primaryKey"`
	Email    string `gorm:"unique; index; not null"`
	Password string `gorm:"not null;"`
}

// FindByEmail 根据邮箱查找单一记录
func (u *UserCredential) FindByEmail(email string) (UserCredential, error) {
	var userCredential UserCredential
	return userCredential, DB.Where("email = ?", email).First(&userCredential).Error
}

// RegisterNewUser 新用户注册的数据库操作
func (u *UserCredential) RegisterNewUser(email, username, passwordHash string) (UserCredential, UserInfo, error) {
	newUserCredential := UserCredential{
		Email:    email,
		Password: passwordHash,
	}
	if err := DB.Create(&newUserCredential).Error; err != nil {
		return UserCredential{}, UserInfo{}, err
	}
	// 同步注册新的 UserInfo
	newUserInfo := UserInfo{
		UserId:   newUserCredential.Id,
		Email:    email,
		Username: username,
	}
	if err := DB.Create(&newUserInfo).Error; err != nil {
		return UserCredential{}, UserInfo{}, err
	}
	return newUserCredential, newUserInfo, nil
}

package dal

type UserCredential struct {
	Id       int64  `gorm:"primaryKey"`
	Username string `gorm:"unique; not null; index"`
	Password string `gorm:"not null;"`
}

func (u *UserCredential) FindById(id int64) (UserCredential, error) {
	var userCredential UserCredential
	return userCredential, DB.First(&userCredential, id).Error
}

func (u *UserCredential) FindByUsername(username string) (UserCredential, error) {
	var userCredential UserCredential
	return userCredential, DB.Where("username = ?", username).First(&userCredential).Error
}

func (u *UserCredential) UpdatePassword(id int64, password string) error {
	var userCredential UserCredential
	if err := DB.First(&userCredential, id).Error; err != nil {
		return err
	}
	userCredential.Password = password
	return DB.Save(&userCredential).Error
}

// RegisterNewUser register a new UserCredential along with a linked UserInfo
func (u *UserCredential) RegisterNewUser(username, passwordHash string) (UserCredential, UserInfo, error) {
	newUserCredential := UserCredential{
		Username: username,
		Password: passwordHash,
	}
	if err := DB.Create(&newUserCredential).Error; err != nil {
		return UserCredential{}, UserInfo{}, err
	}
	// linked UserInfo
	newUserInfo := UserInfo{
		UserId:   newUserCredential.Id,
		Username: username,
	}
	if err := newUserInfo.Create(&newUserInfo); err != nil {
		return UserCredential{}, UserInfo{}, err
	}
	return newUserCredential, newUserInfo, nil
}

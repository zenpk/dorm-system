package dal

type UserCredential struct {
	Id       uint64 `gorm:"primaryKey"`
	Username string `gorm:"unique; not null; index"`
	Password string `gorm:"not null;"`
}

func (u *UserCredential) FindById(id uint64) (*UserCredential, error) {
	userCredential := new(UserCredential)
	return userCredential, DB.First(&userCredential, id).Error
}

func (u *UserCredential) FindByUsername(username string) (*UserCredential, error) {
	userCredential := new(UserCredential)
	return userCredential, DB.Where("username = ?", username).First(&userCredential).Error
}

func (u *UserCredential) Update(credential *UserCredential) error {
	return DB.Save(&credential).Error
}

// RegisterNewUser register a new UserCredential along with a linked UserInfo
func (u *UserCredential) RegisterNewUser(username, passwordHash string) (*UserCredential, *UserInfo, error) {
	newUserCredential := &UserCredential{
		Username: username,
		Password: passwordHash,
	}
	if err := DB.Create(&newUserCredential).Error; err != nil {
		return nil, nil, err
	}
	// linked UserInfo
	newUserInfo := &UserInfo{
		UserId:   newUserCredential.Id,
		Username: username,
	}
	if err := newUserInfo.Create(newUserInfo); err != nil {
		return nil, nil, err
	}
	return newUserCredential, newUserInfo, nil
}

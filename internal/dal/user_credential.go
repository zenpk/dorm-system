package dal

import "context"

type UserCredential struct {
	Id       uint64 `gorm:"primaryKey"`
	Username string `gorm:"unique; not null; index"`
	Password string `gorm:"not null;"`
}

func (u *UserCredential) FindById(ctx context.Context, id uint64) (*UserCredential, error) {
	userCredential := new(UserCredential)
	return userCredential, DB.WithContext(ctx).First(&userCredential, id).Error
}

func (u *UserCredential) FindByUsername(ctx context.Context, username string) (*UserCredential, error) {
	userCredential := new(UserCredential)
	return userCredential, DB.WithContext(ctx).Where("username = ?", username).First(&userCredential).Error
}

func (u *UserCredential) Update(ctx context.Context, credential *UserCredential) error {
	return DB.WithContext(ctx).Save(&credential).Error
}

// RegisterNewUser register a new UserCredential along with a linked UserInfo
func (u *UserCredential) RegisterNewUser(ctx context.Context, username, passwordHash string) (*UserCredential, *UserInfo, error) {
	newUserCredential := &UserCredential{
		Username: username,
		Password: passwordHash,
	}
	if err := DB.WithContext(ctx).Create(&newUserCredential).Error; err != nil {
		return nil, nil, err
	}
	// linked UserInfo
	newUserInfo := &UserInfo{
		UserCredentialId: newUserCredential.Id,
		Username:         username,
	}
	if err := newUserInfo.Create(ctx, newUserInfo); err != nil {
		return nil, nil, err
	}
	return newUserCredential, newUserInfo, nil
}

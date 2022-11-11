package dal

import "context"

type Account struct {
	Id       uint64 `gorm:"primaryKey"`
	Uid      uint64 `gorm:"not null; index"`
	Username string `gorm:"not null; unique; index"`
	Password string `gorm:"not null"`
	Deleted  bool   `gorm:"not null; default:0; index"`
}

func (a *Account) FindById(ctx context.Context, id uint64) (*Account, error) {
	userCredential := new(Account)
	return userCredential, DB.WithContext(ctx).First(&userCredential, id).Error
}

func (a *Account) FindByUsername(ctx context.Context, username string) (*Account, error) {
	userCredential := new(Account)
	return userCredential, DB.WithContext(ctx).Where("username = ?", username).First(&userCredential).Error
}

func (a *Account) Update(ctx context.Context, credential *Account) error {
	return DB.WithContext(ctx).Save(&credential).Error
}

// RegisterNewUser register a new Account along with a linked User
func (a *Account) RegisterNewUser(ctx context.Context, username, passwordHash string) (*Account, *User, error) {
	// create User and get uid
	user := &User{}
	if err := user.Create(ctx, user); err != nil {
		return nil, nil, err
	}
	account := &Account{
		Uid:      user.Id,
		Username: username,
		Password: passwordHash,
	}
	if err := DB.WithContext(ctx).Create(&account).Error; err != nil {
		return nil, nil, err
	}
	return account, user, nil
}

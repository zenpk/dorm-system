package dal

import "context"

type Account struct {
	Id       uint64 `gorm:"primaryKey"`
	UserId   uint64 `gorm:"not null; index"`
	Username string `gorm:"not null; unique; index"`
	Password string `gorm:"not null"`
	Deleted  bool   `gorm:"not null; default:0; index"`
}

func (a *Account) FindById(ctx context.Context, id uint64) (*Account, error) {
	account := new(Account)
	return account, DB.WithContext(ctx).First(&account, id).Error
}

func (a *Account) FindByUserId(ctx context.Context, id uint64) (*Account, error) {
	account := new(Account)
	return account, DB.WithContext(ctx).Where("user_id = ?", id).First(&account).Error
}

func (a *Account) FindByUsername(ctx context.Context, username string) (*Account, error) {
	account := new(Account)
	return account, DB.WithContext(ctx).Where("username = ?", username).First(&account).Error
}

func (a *Account) Update(ctx context.Context, account *Account) error {
	return DB.WithContext(ctx).Save(&account).Error
}

// RegisterNewUser register a new Account along with a linked User
func (a *Account) RegisterNewUser(ctx context.Context, username, passwordHash string) (*User, error) {
	// create User and get uid
	user := &User{}
	if err := user.Create(ctx, user); err != nil {
		return nil, err
	}
	account := &Account{
		UserId:   user.Id,
		Username: username,
		Password: passwordHash,
	}
	if err := DB.WithContext(ctx).Create(&account).Error; err != nil {
		return nil, err
	}
	return user, nil
}

package dal

import (
	"context"
	"gorm.io/gorm"
)

type Account struct {
	Id       uint64 `gorm:"primaryKey"`
	UserId   uint64 `gorm:"not null"`
	Username string `gorm:"not null; unique"`
	Password string `gorm:"not null"`
	Deleted  gorm.DeletedAt
}

func (a Account) FindById(ctx context.Context, id uint64) (account *Account, err error) {
	return account, DB.WithContext(ctx).Take(&account, id).Error
}

func (a Account) FindByUserId(ctx context.Context, id uint64) (account *Account, err error) {
	return account, DB.WithContext(ctx).Where("user_id = ?", id).Take(&account).Error
}

func (a Account) FindByUsername(ctx context.Context, username string) (account *Account, err error) {
	return account, DB.WithContext(ctx).Where("username = ?", username).Take(&account).Error
}

func (a Account) Update(ctx context.Context, account *Account) error {
	return DB.WithContext(ctx).Save(&account).Error
}

// RegisterNewUser register a new Account along with a linked User, using transaction
func (a Account) RegisterNewUser(ctx context.Context, username, passwordHash string) (user *User, err error) {
	err = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// create User and get uid
		if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
			return err
		}
		account := &Account{
			UserId:   user.Id,
			Username: username,
			Password: passwordHash,
		}
		if err := DB.WithContext(ctx).Create(&account).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, err
}

package dal

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/util"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	Id           uint64    `gorm:"primaryKey" json:"-"`
	RefreshToken string    `gorm:"not null" json:"refreshToken,omitempty"`
	UserId       uint64    `gorm:"not null" json:"-"`
	CreateTime   time.Time `gorm:"not null" json:"createTime,omitempty"`
	ExpTime      time.Time `gorm:"not null" json:"expTime,omitempty"`
	Deleted      gorm.DeletedAt
}

func (t Token) GenNew(ctx context.Context, id uint64) (refreshToken string, err error) {
	node, err := snowflake.NewNode(viper.GetInt64("snowflake.node"))
	if err != nil {
		return "", err
	}
	createTime := util.GetUTCTime()
	age := time.Duration(viper.GetInt64("refresh_token.age_hour")) * time.Hour
	expTime := createTime.Add(age)
	snowflakeId := node.Generate()
	token := &Token{
		RefreshToken: snowflakeId.Base64(),
		UserId:       id,
		CreateTime:   createTime,
		ExpTime:      expTime,
	}
	return token.RefreshToken, DB.WithContext(ctx).Create(&token).Error
}

func (t Token) FindByRefreshToken(ctx context.Context, refreshToken string) (token *Token, err error) {
	return token, DB.WithContext(ctx).Where("refresh_token = ?", refreshToken).Take(&token).Error
}

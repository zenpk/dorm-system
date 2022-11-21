package dal

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
)

type Team struct {
	Id      uint64 `gorm:"primaryKey"`
	Code    string `gorm:"not null; unique; index"`
	Gender  string `gorm:"size:10; not null"`
	OwnerId uint64 `gorm:"not null; unique; index"`
	Deleted bool   `gorm:"not null; default:0; index"`
}

func (t *Team) GenNew(ctx context.Context, owner *User) (*Team, error) {
	node, err := snowflake.NewNode(viper.GetInt64("snowflake.node"))
	if err != nil {
		return nil, err
	}
	snowflakeId := node.Generate()
	team := &Team{
		Code:    snowflakeId.Base64(),
		Gender:  owner.Gender,
		OwnerId: owner.Id,
	}
	return team, DB.WithContext(ctx).Create(&team).Error
}

package dal

import (
	"context"
	"errors"
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Team struct {
	Id      uint64         `gorm:"primaryKey"`
	Code    string         `gorm:"not null; unique; index"`
	Gender  string         `gorm:"size:10; not null"`
	OwnerId uint64         `gorm:"not null; unique; index"`
	Deleted gorm.DeletedAt `gorm:"index"`
}

func (t Team) FindByOwnerId(ctx context.Context, id uint64) (*Team, error) {
	team := new(Team)
	return team, DB.WithContext(ctx).Where("owner_id = ?", id).Take(&team).Error
}

// FindByInnerJoinUserId only finds team member, for owner please use FindByOwnerId
func (t Team) FindByInnerJoinUserId(ctx context.Context, id uint64) (*Team, error) {
	team := new(Team)
	// alternative: use Raw, cons: have to explicitly set WHERE `deleted` IS NULL for `teams`
	//DB.WithContext(ctx).Raw("SELECT t.* FROM `teams` t INNER JOIN `team_users` tu ON tu.`team_id` = t.`id` WHERE t.`deleted` IS NULL AND tu.`deleted` IS NULL AND tu.`user_id` = ?", id).Take(&team)
	return team, DB.WithContext(ctx).Select("`teams`.*").Joins("INNER JOIN (SELECT * FROM `team_users` WHERE `deleted` IS NULL AND `user_id` = ?) tu ON tu.`team_id` = `teams`.`id`", id).Take(&team).Error
}

// CheckIfHasTeam checks if a user is a team owner OR a team member
func (t Team) CheckIfHasTeam(ctx context.Context, userId uint64) (*Team, error) {
	team := new(Team)
	if err := DB.WithContext(ctx).Where("owner_id = ?", userId).Take(&team).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // not owner, but can be a member
			team, err = t.FindByInnerJoinUserId(ctx, userId)
			if err != nil { // include ErrRecordNotFound
				return nil, err
			}
		} else { // other errors
			return nil, err
		}
	}
	return team, nil
}

func (t Team) FindByCode(ctx context.Context, code string) (team *Team, err error) {
	return team, DB.WithContext(ctx).Where("code = ?", code).Take(&team).Error
}

func (t Team) GenNew(ctx context.Context, owner *User) (team *Team, err error) {
	node, err := snowflake.NewNode(viper.GetInt64("snowflake.node"))
	if err != nil {
		return nil, err
	}
	snowflakeId := node.Generate()
	team = &Team{
		Code:    snowflakeId.Base64(),
		Gender:  owner.Gender,
		OwnerId: owner.Id,
	}
	return team, DB.WithContext(ctx).Create(&team).Error
}

func (t Team) TransSetNewOwner(ctx context.Context, team *Team, rel *TeamUser) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// update owner id
		team.OwnerId = rel.UserId
		if err := t.Update(ctx, team); err != nil {
			return err
		}
		// delete member record
		if err := Table.TeamUser.Delete(ctx, rel); err != nil {
			return err
		}
		return nil
	})
}

func (t Team) Update(ctx context.Context, team *Team) error {
	return DB.WithContext(ctx).Save(&team).Error
}

func (t Team) Delete(ctx context.Context, team *Team) error {
	return DB.WithContext(ctx).Delete(&team).Error
}

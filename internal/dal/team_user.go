package dal

import (
	"context"
	"gorm.io/gorm"
)

type TeamUser struct {
	Id      uint64         `gorm:"primaryKey"`
	TeamId  uint64         `gorm:"not null; index"`
	UserId  uint64         `gorm:"not null; index"`
	Deleted gorm.DeletedAt `gorm:"index"`
}

func (t *TeamUser) CntTeamMember(ctx context.Context, teamId uint64) (cnt uint64, err error) {
	// + 1 because owner
	return cnt + 1, DB.WithContext(ctx).Model(&TeamUser{}).Select("COUNT(*)").Where("team_id = ?", teamId).Error
}

func (t *TeamUser) PluckAllUserIdsByTeamId(ctx context.Context, teamId uint64) (userIds []uint64, err error) {
	return userIds, DB.WithContext(ctx).Model(&TeamUser{}).Select("user_id").Where("team_id = ?", teamId).Scan(&userIds).Error
}

func (t *TeamUser) Create(ctx context.Context, teamUser *TeamUser) error {
	return DB.WithContext(ctx).Create(&teamUser).Error
}

func (t *TeamUser) Delete(ctx context.Context, teamUser *TeamUser) error {
	return DB.WithContext(ctx).Delete(&teamUser).Error
}

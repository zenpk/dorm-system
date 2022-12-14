package dal

import (
	"context"
	"gorm.io/gorm"
)

type TeamUser struct {
	Id      uint64 `gorm:"primaryKey"`
	TeamId  uint64 `gorm:"not null"`
	UserId  uint64 `gorm:"not null"`
	Deleted gorm.DeletedAt
}

func (t TeamUser) FindByTeamIdAndUserId(ctx context.Context, teamId, userId uint64) (rel *TeamUser, err error) {
	return rel, DB.WithContext(ctx).Where("team_id = ? AND user_id = ?", teamId, userId).Take(&rel).Error
}

func (t TeamUser) CntTeamMember(ctx context.Context, teamId uint64) (cnt uint64, err error) {
	// + 1 because owner
	return cnt + 1, DB.WithContext(ctx).Model(&TeamUser{}).Select("COUNT(*)").Where("team_id = ?", teamId).Error
}

func (t TeamUser) PluckAllUserIdsByTeamId(ctx context.Context, teamId uint64) (userIds []uint64, err error) {
	return userIds, DB.WithContext(ctx).Model(&TeamUser{}).Select("user_id").Where("team_id = ?", teamId).Order("id").Scan(&userIds).Error
}

func (t TeamUser) Create(ctx context.Context, teamUser *TeamUser) error {
	return DB.WithContext(ctx).Create(&teamUser).Error
}

func (t TeamUser) Delete(ctx context.Context, teamUser *TeamUser) error {
	return DB.WithContext(ctx).Delete(&teamUser).Error
}

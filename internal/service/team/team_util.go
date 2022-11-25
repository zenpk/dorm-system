package team

import (
	"context"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/service/user"
)

// getOwnerAndMembers from team, return
func (s Server) getOwnerAndMembers(ctx context.Context, team *dal.Team) (owner *user.UserInfo, members []*user.UserInfo, err error) {
	ownerUser, err := dal.Table.User.FindById(ctx, team.OwnerId)
	if err != nil {
		return nil, nil, err
	}
	owner = &user.UserInfo{
		Id:         ownerUser.Id,
		StudentNum: ownerUser.StudentNum,
		Name:       ownerUser.Name,
		Gender:     ownerUser.Gender,
	}
	ids, err := dal.Table.TeamUser.PluckAllUserIdsByTeamId(ctx, team.Id)
	if err != nil {
		return nil, nil, err
	}
	memberUsers, err := dal.Table.User.FindAllByIds(ctx, ids)
	if err != nil {
		return nil, nil, err
	}
	for _, mu := range memberUsers {
		member := &user.UserInfo{
			Id:         mu.Id,
			StudentNum: mu.StudentNum,
			Name:       mu.Name,
			Gender:     mu.Gender,
		}
		members = append(members, member)
	}
	return owner, members, nil
}

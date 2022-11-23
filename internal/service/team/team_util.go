package team

import (
	"context"
	"github.com/zenpk/dorm-system/internal/dal"
)

// getOwnerAndMembers from team, return
func (s *Server) getOwnerAndMembers(ctx context.Context, team *dal.Team) (owner *UserInfo, members []*UserInfo, err error) {
	ownerUser, err := dal.Table.User.FindById(ctx, team.OwnerId)
	if err != nil {
		return nil, nil, err
	}
	owner.StudentNum = ownerUser.StudentNum
	owner.Name = ownerUser.Name
	ids, err := dal.Table.TeamUser.PluckAllUserIdsByTeamId(ctx, team.Id)
	if err != nil {
		return nil, nil, err
	}
	memberUsers, err := dal.Table.User.FindAllByIds(ctx, ids)
	if err != nil {
		return nil, nil, err
	}
	for _, mu := range memberUsers {
		member := &UserInfo{
			StudentNum: mu.StudentNum,
			Name:       mu.Name,
		}
		members = append(members, member)
	}
	return owner, members, nil
}

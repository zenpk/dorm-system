package team

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/service/common"
	"github.com/zenpk/dorm-system/pkg/ep"
	"gorm.io/gorm"
)

type Server struct {
	Config *viper.Viper
	UnimplementedTeamServer
}

func (s Server) Create(ctx context.Context, req *CreateRequest) (*CreateReply, error) {
	// check if user already has a team
	team, err := dal.Table.Team.CheckIfHasTeam(ctx, req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// user not in a team, continue
		user, err := dal.Table.User.FindById(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		newTeam, err := dal.Table.Team.GenNew(ctx, user)
		if err != nil {
			return nil, err
		}
		resp := &CreateReply{
			Err: &common.CommonResponse{
				Code: ep.ErrOK.Code,
				Msg:  ep.ErrOK.Msg,
			},
			Code: newTeam.Code,
		}
		return resp, nil
	} else if err != nil { // other errors
		return nil, err
	}
	// user already has a team
	resp := &CreateReply{
		Err: &common.CommonResponse{
			Code: ep.ErrDuplicatedRecord.Code,
			Msg:  "user already has a team",
		},
		Code: team.Code,
	}
	return resp, nil
}

func (s Server) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	team, err := dal.Table.Team.CheckIfHasTeam(ctx, req.UserId)
	if err != nil { // include ErrRecordNotFound
		return nil, err
	}
	owner, members, err := s.getOwnerAndMembers(ctx, team)
	if err != nil {
		return nil, err
	}
	resp := &GetReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  ep.ErrOK.Msg,
		},
		Team: &TeamInfo{Id: team.Id,
			Code:    team.Code,
			Gender:  team.Gender,
			Owner:   owner,
			Members: members,
		},
	}
	return resp, nil
}

func (s Server) Join(ctx context.Context, req *JoinRequest) (*JoinReply, error) {
	// check if user already has a team
	_, err := dal.Table.Team.CheckIfHasTeam(ctx, req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// user not in a team, continue
		user, err := dal.Table.User.FindById(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		targetTeam, err := dal.Table.Team.FindByCode(ctx, req.Code)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			resp := &JoinReply{
				Err: &common.CommonResponse{
					Code: ep.ErrInputBody.Code,
					Msg:  "cannot find a team with given code",
				},
			}
			return resp, nil
		} else if err != nil { // other errors
			return nil, err
		}
		// check gender
		if user.Gender != targetTeam.Gender {
			resp := &JoinReply{
				Err: &common.CommonResponse{
					Code: ep.ErrLogic.Code,
					Msg:  "cannot join a team with a different gender",
				},
			}
			return resp, nil
		}
		// check member number limit
		nowCnt, err := dal.Table.TeamUser.CntTeamMember(ctx, targetTeam.Id)
		if nowCnt >= s.Config.GetUint64("max_member_lim") {
			resp := &JoinReply{
				Err: &common.CommonResponse{
					Code: ep.ErrLogic.Code,
					Msg:  "target team is already full",
				},
			}
			return resp, nil
		}
		// create join relation
		rel := &dal.TeamUser{
			TeamId: targetTeam.Id,
			UserId: user.Id,
		}
		if err := dal.Table.TeamUser.Create(ctx, rel); err != nil {
			return nil, err
		}
		resp := &JoinReply{
			Err: &common.CommonResponse{
				Code: ep.ErrOK.Code,
				Msg:  ep.ErrOK.Msg,
			},
		}
		return resp, nil
	} else if err != nil { // other errors
		return nil, err
	}
	// user already has a team
	resp := &JoinReply{
		Err: &common.CommonResponse{
			Code: ep.ErrDuplicatedRecord.Code,
			Msg:  "user already has a team",
		},
	}
	return resp, nil
}

func (s Server) Leave(ctx context.Context, req *LeaveRequest) (*LeaveReply, error) {
	// check if user already has a team
	team, err := dal.Table.Team.CheckIfHasTeam(ctx, req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp := &LeaveReply{
			Err: &common.CommonResponse{
				Code: ep.ErrNoRecord.Code,
				Msg:  "user don't have a team",
			},
		}
		return resp, nil
	} else if err != nil { // other errors
		return nil, err
	}
	if team.OwnerId != req.UserId { // not the owner
		rel, err := dal.Table.TeamUser.FindByTeamIdAndUserId(ctx, team.Id, req.UserId)
		if err != nil {
			return nil, err
		}
		if err := dal.Table.TeamUser.Delete(ctx, rel); err != nil {
			return nil, err
		}
		resp := &LeaveReply{
			Err: &common.CommonResponse{
				Code: ep.ErrOK.Code,
				Msg:  "successfully left the team",
			},
		}
		return resp, nil
	}
	// user is the owner, if user is the only one in the team
	// then delete the whole team
	memberIds, err := dal.Table.TeamUser.PluckAllUserIdsByTeamId(ctx, team.Id)
	if len(memberIds) > 0 { // if the team contains other users
		resp := &LeaveReply{
			Err: &common.CommonResponse{
				Code: ep.ErrLogic.Code,
				Msg:  "you need to transfer the ownership before leaving",
			},
		}
		return resp, nil
	}
	if err := dal.Table.Team.Delete(ctx, team); err != nil {
		return nil, err
	}
	resp := &LeaveReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "successfully left the team, the team is deleted",
		},
	}
	return resp, nil
}

func (s Server) Transfer(ctx context.Context, req *TransferRequest) (*TransferReply, error) {
	// check if user already has a team
	team, err := dal.Table.Team.CheckIfHasTeam(ctx, req.OldOwnerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp := &TransferReply{
			Err: &common.CommonResponse{
				Code: ep.ErrNoRecord.Code,
				Msg:  "user don't have a team",
			},
		}
		return resp, nil
	} else if err != nil { // other errors
		return nil, err
	}
	if team.OwnerId != req.OldOwnerId { // not the owner
		resp := &TransferReply{
			Err: &common.CommonResponse{
				Code: ep.ErrNoPermission.Code,
				Msg:  "you're not the team's owner",
			},
		}
		return resp, nil
	}
	// check if new owner is a member of the team
	rel, err := dal.Table.TeamUser.FindByTeamIdAndUserId(ctx, team.Id, req.NewOwnerId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		resp := &TransferReply{
			Err: &common.CommonResponse{
				Code: ep.ErrNoRecord.Code,
				Msg:  "new owner doesn't belong to the team",
			},
		}
		return resp, nil
	} else if err != nil {
		return nil, err
	}
	// transfer the ownership
	// start a transaction to do the operation
	if err := dal.Table.Team.TransSetNewOwner(ctx, team, rel); err != nil {
		return nil, err
	}
	resp := &TransferReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "successfully transferred the ownership",
		},
	}
	return resp, nil
}

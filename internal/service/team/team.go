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

func (s *Server) Create(ctx context.Context, req *CreateGetRequest) (*CreateReply, error) {
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
			Resp: &common.CommonResponse{
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
		Resp: &common.CommonResponse{
			Code: ep.ErrDuplicatedRecord.Code,
			Msg:  "user already has a team",
		},
		Code: team.Code,
	}
	return resp, nil
}

func (s *Server) Get(ctx context.Context, req *CreateGetRequest) (*GetReply, error) {
	team, err := dal.Table.Team.CheckIfHasTeam(ctx, req.UserId)
	if err != nil { // include ErrRecordNotFound
		return nil, err
	}
	owner, members, err := s.getOwnerAndMembers(ctx, team)
	if err != nil {
		return nil, err
	}
	resp := &GetReply{
		Resp: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  ep.ErrOK.Msg,
		},
		Id:      team.Id,
		Code:    team.Code,
		Gender:  team.Gender,
		Owner:   owner,
		Members: members,
	}
	return resp, nil
}

func (s *Server) Join(ctx context.Context, req *JoinRequest) (*JoinReply, error) {
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
				Resp: &common.CommonResponse{
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
				Resp: &common.CommonResponse{
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
				Resp: &common.CommonResponse{
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
			Resp: &common.CommonResponse{
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
		Resp: &common.CommonResponse{
			Code: ep.ErrDuplicatedRecord.Code,
			Msg:  "user already has a team",
		},
	}
	return resp, nil
}

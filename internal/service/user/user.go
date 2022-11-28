package user

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/service/common"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/ep"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Server struct {
	Config *viper.Viper
	UnimplementedUserServer
}

func (s Server) Register(ctx context.Context, req *RegisterLoginRequest) (*UserReply, error) {
	// username duplication check
	_, err := dal.Table.Account.FindByUsername(ctx, req.Username)
	if err == nil { // user already exists
		errPack := ep.ErrDuplicatedRecord
		errPack.Msg = "user already exists"
		return nil, errPack
	} else if !errors.Is(err, gorm.ErrRecordNotFound) { // something went wrong
		return nil, err
	}
	// no duplication, register
	passwordHash, err := util.BCryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user, err := dal.Table.Account.RegisterNewUser(ctx, req.Username, passwordHash)
	if err != nil {
		return nil, err
	}
	resp := &UserReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "successfully registered",
		},
		UserId: user.Id,
	}
	return resp, nil
}

func (s Server) Login(ctx context.Context, req *RegisterLoginRequest) (*UserReply, error) {
	account, err := dal.Table.Account.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	passwordHashByte := []byte(account.Password)
	passwordByte := []byte(req.Password)
	if err := bcrypt.CompareHashAndPassword(passwordHashByte, passwordByte); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		errPack := ep.ErrInputBody
		errPack.Msg = "wrong password"
		return nil, errPack
	}
	user, err := dal.Table.User.FindById(ctx, account.UserId)
	if err != nil {
		return nil, ep.ErrNoRecord
	}
	resp := &UserReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "successfully logged in",
		},
		UserId: user.Id,
	}
	return resp, nil
}

func (s Server) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	user, err := dal.Table.User.FindById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	resp := &GetReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "successfully got user information",
		},
		User: &UserInfo{
			Id:         user.Id,
			StudentNum: user.StudentNum,
			Name:       user.Name,
			Gender:     user.Gender,
		},
	}
	return resp, nil
}

func (s Server) Edit(ctx context.Context, req *EditRequest) (*EditReply, error) {
	user, err := dal.Table.User.FindById(ctx, req.User.Id)
	if err != nil {
		return nil, err
	}
	if req.User.Gender != user.Gender {
		// user need to leave team to edit his/her gender
		_, err := dal.Table.Team.CheckIfHasTeam(ctx, user.Id)
		if err == nil { // user has team
			resp := &EditReply{
				Err: &common.CommonResponse{
					Code: ep.ErrLogic.Code,
					Msg:  "you must leave your team before editing gender",
				},
			}
			return resp, nil
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // other errors
			return nil, err
		}
	}
	// update
	user.StudentNum = req.User.StudentNum
	user.Name = req.User.Name
	user.Gender = req.User.Gender
	if err := dal.Table.User.Update(ctx, user); err != nil {
		return nil, err
	}
	resp := &EditReply{
		Err: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  "successfully updated user information",
		},
	}
	return resp, nil
}

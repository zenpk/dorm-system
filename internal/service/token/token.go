package token

import (
	"context"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/service/common"
	"github.com/zenpk/dorm-system/pkg/ep"
	"github.com/zenpk/dorm-system/pkg/jwt"
)

type Server struct {
	Config *viper.Viper
	UnimplementedTokenServer
}

func (s *Server) GenAllToken(ctx context.Context, req *GenAllTokenRequest) (*TokenReply, error) {
	user, err := dal.Table.User.FindById(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	account, err := dal.Table.Account.FindByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	claims := &jwt.MyCustomClaims{
		UserId:   user.Id,
		Username: account.Username,
		Role:     user.Role,
	}
	accessToken, err := jwt.GenToken(claims)
	if err != nil {
		return nil, err
	}
	refreshToken, err := dal.Table.Token.GenNew(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	resp := &TokenReply{
		Resp: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  ep.ErrOK.Msg,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return resp, nil
}

func (s *Server) GenAccessToken(ctx context.Context, req *GenAccessTokenRequest) (*TokenReply, error) {
	token, err := dal.Table.Token.FindByRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}
	user, err := dal.Table.User.FindById(ctx, token.UserId)
	if err != nil {
		return nil, err
	}
	account, err := dal.Table.Account.FindByUserId(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	claims := &jwt.MyCustomClaims{
		UserId:   user.Id,
		Username: account.Username,
		Role:     user.Role,
	}
	accessToken, err := jwt.GenToken(claims)
	if err != nil {
		return nil, err
	}
	refreshToken, err := dal.Table.Token.GenNew(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	resp := &TokenReply{
		Resp: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  ep.ErrOK.Msg,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return resp, nil
}

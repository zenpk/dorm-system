package user

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/service/common"
	"github.com/zenpk/dorm-system/pkg/ep"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Server struct {
	Config *viper.Viper
	UnimplementedUserServer
}

func (s *Server) Register(ctx context.Context, req *RegisterLoginRequest) (*UserReply, error) {
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
	passwordHash, err := bCryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	user, err := dal.Table.Account.RegisterNewUser(ctx, req.Username, passwordHash)
	if err != nil {
		return nil, err
	}
	resp := &UserReply{
		Resp: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  ep.ErrOK.Msg,
		},
		UserId: user.Id,
	}
	return resp, nil
}

func (s *Server) Login(ctx context.Context, req *RegisterLoginRequest) (*UserReply, error) {
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
		Resp: &common.CommonResponse{
			Code: ep.ErrOK.Code,
			Msg:  ep.ErrOK.Msg,
		},
		UserId: user.Id,
	}
	return resp, nil
}

//// GetMyInfo get User based on the id in Cookie
//func (*User) GetMyInfo(c *gin.Context) {
//	idStr, err := cookie.GetUserId(c)
//	if err != nil {
//		c.JSON(http.StatusOK, dto.GetUserInfoResp{
//			CommonResp: dto.CommonResp{
//				Code: eh.Preset.CodeTokenError,
//				Msg:  "you're not logged in",
//			},
//		})
//		return
//	}
//	id := util.ParseU64(idStr)
//	var userInfo *dal.User
//	userInfo, err = userInfo.FindById(id)
//	errHandler := eh.JSONHandler{C: c, V: dto.GetUserInfoResp{}}
//	if err != nil {
//		errHandler.Handle(err)
//		return
//	}
//	c.JSON(http.StatusOK, dto.GetUserInfoResp{
//		CommonResp: dto.CommonResp{
//			Code: eh.Preset.CodeOK,
//			Msg:  "success",
//		},
//		User: userInfo,
//	})
//}

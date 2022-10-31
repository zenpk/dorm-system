package user

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/service/common"
	"github.com/zenpk/dorm-system/pkg/ep"
	"github.com/zenpk/dorm-system/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Server struct {
	Config *viper.Viper
	UnimplementedUserServer
}

func (s *Server) Register(ctx context.Context, req *RegisterLoginRequest) (*UserReply, error) {
	// username duplication check
	table := new(dal.UserCredential)
	_, err := table.FindByUsername(ctx, req.Username)
	if err == nil { // user already exists
		errPack := ep.ErrDuplicateRecord
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
	newUserCredential, _, err := table.RegisterNewUser(ctx, req.Username, passwordHash)
	if err != nil {
		return nil, err
	}
	// generate JWT
	token, err := jwt.GenToken(newUserCredential.Id, req.Username)
	if err != nil {
		return nil, err
	}
	resp := new(UserReply)
	resp.Resp = &common.CommonResponse{
		Code: 0,
		Msg:  "successfully registered",
	}
	resp.Token = token
	resp.UserId = newUserCredential.Id
	resp.Username = req.Username
	return resp, nil
}

func (s *Server) Login(ctx context.Context, req *RegisterLoginRequest) (*UserReply, error) {
	userCredential, err := dal.Table.UserCredential.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	passwordHashByte := []byte(userCredential.Password)
	passwordByte := []byte(req.Password)
	if err := bcrypt.CompareHashAndPassword(passwordHashByte, passwordByte); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		errPack := ep.ErrInputBody
		errPack.Msg = "wrong password"
		return nil, errPack
	}
	token, err := jwt.GenToken(userCredential.Id, req.Username)
	if err != nil {
		return nil, err
	}
	resp := new(UserReply)
	resp.Resp = &common.CommonResponse{
		Code: ep.ErrOK.Code,
		Msg:  "successfully logged in",
	}
	resp.Token = token
	resp.UserId = userCredential.Id
	resp.Username = req.Username
	return resp, nil
}

//// GetMyInfo get UserInfo based on the id in Cookie
//func (*UserInfo) GetMyInfo(c *gin.Context) {
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
//	var userInfo *dal.UserInfo
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
//		UserInfo: userInfo,
//	})
//}

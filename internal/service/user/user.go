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
	_, err := table.FindByUsername(req.Username)
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
	newUserCredential, _, err := table.RegisterNewUser(req.Username, passwordHash)
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

//// Login 用户登录
//func (s *Server) Login(ctx context.Context, req *RegisterLoginRequest) (*UserReply, error) {
//	var req dto.LoginReq
//	packer := ep.Packer{V: dto.RegisterLoginResp{}}
//	if err := c.ShouldBindJSON(&req); err != nil {
//		Response(c, packer.PackWithError(err))
//		return
//	}
//	var credentialTable *dal.UserCredential
//	userCredential, err := credentialTable.FindByEmail(req.Email)
//	if errors.Is(err, gorm.ErrRecordNotFound) { // 没有查到该邮箱对应的记录
//		Response(c, dto.RegisterLoginResp{
//			CommonResp: dto.CommonResp{
//				SuccessCode: false,
//				Code:        ep.ErrNoRecord.Code,
//				Msg:         "用户不存在",
//			},
//			Status: 1,
//		})
//		return
//	} else if err != nil { // 出现其他错误
//		Response(c, packer.PackWithError(err))
//		return
//	}
//	// 有对应的用户名，则比较密码
//	passwordHashByte := []byte(userCredential.Password)
//	passwordByte := []byte(req.Password)
//	// 检查密码是否正确
//	if err := bcrypt.CompareHashAndPassword(passwordHashByte, passwordByte); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
//		Response(c, dto.RegisterLoginResp{
//			CommonResp: dto.CommonResp{
//				SuccessCode: false,
//				Code:        ep.ErrInputBody.Code,
//				Msg:         "密码错误",
//			},
//			Status: 2,
//		})
//		return
//	}
//	var infoTable *dal.UserInfo
//	userInfo, err := infoTable.FindByUserCredentialId(userCredential.Id)
//	if err != nil {
//		Response(c, packer.PackWithError(err))
//		return
//	}
//	// 生成 JWT
//	token, err := jwt.GenToken(userCredential.Id, userInfo.Username)
//	if err != nil {
//		Response(c, packer.PackWithError(err))
//		return
//	}
//	// 写入 Cookie
//	cookie.SetToken(c, token)
//	cookie.SetUserId(c, strconv.FormatUint(userCredential.Id, 10))
//	cookie.SetUsername(c, userInfo.Username)
//	Response(c, dto.RegisterLoginResp{
//		CommonResp: dto.CommonResp{
//			SuccessCode: true,
//			Code:        ep.ErrOK.Code,
//			Msg:         "登录成功",
//		},
//		Status: 0,
//		Data:   userInfo,
//	})
//}
//
//// Logout 用户登出
//func (s *Server) Logout(ctx context.Context, req *RegisterLoginRequest) (*UserReply, error) {
//	packer := ep.Packer{V: dto.RegisterLoginResp{}}
//	if _, err := cookie.GetUserId(c); err != nil {
//		Response(c, packer.PackWithError(err))
//		return
//	}
//	cookie.ClearAllUserInfos(c)
//	Response(c, dto.RegisterLoginResp{
//		CommonResp: dto.CommonResp{
//			SuccessCode: true,
//			Code:        ep.ErrOK.Code,
//			Msg:         "登出成功",
//		},
//		Status: 0,
//		Data:   nil,
//	})
//}

// bCryptPassword 密码 BCrypt 加密
func bCryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

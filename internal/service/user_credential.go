package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/eh"
	"github.com/zenpk/dorm-system/internal/middleware"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserCredential struct{}

func (u *UserCredential) Register(c *gin.Context) {
	var req dto.RegisterLoginReq
	errHandler := eh.User{C: c}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	// check password length
	if len(req.Password) < viper.GetInt("auth.password_length") {
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			Code:     eh.CodeInputError,
			Msg:      "password too short",
			UserInfo: dal.UserInfo{},
		})
		return
	}
	// username duplication check
	var userCredentialDal dal.UserCredential
	_, err := userCredentialDal.FindByUsername(req.Username)
	if err == nil { // record found
		zap.Logger.Warn("user already exists")
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			Code:     eh.CodeDatabaseError,
			Msg:      "user already exists",
			UserInfo: dal.UserInfo{},
		})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) { // not "Not Found" error
		errHandler.RegisterLoginErr(err)
		return
	}
	// no duplication
	passwordHash, err := bCryptPassword(req.Password)
	if err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	newUserCredential, newUserInfo, err := userCredentialDal.RegisterNewUser(req.Username, passwordHash)
	if err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	// generate JWT
	token, err := middleware.GenToken(newUserCredential.Id, req.Username)
	if err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	cookie.SetToken(c, token)
	cookie.SetUserId(c, strconv.FormatInt(newUserCredential.Id, 10))
	cookie.SetUsername(c, req.Username)
	c.JSON(http.StatusOK, dto.RegisterLoginResp{
		Code:     eh.CodeOK,
		Msg:      "successfully registered",
		UserInfo: newUserInfo,
	})
}

func (u *UserCredential) Login(c *gin.Context) {
	var req dto.RegisterLoginReq
	errHandler := eh.User{C: c}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	var userCredentialDal dal.UserCredential
	userCredential, err := userCredentialDal.FindByUsername(req.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) { // no record
		zap.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			Code:     eh.CodeDatabaseError,
			Msg:      "user not exits",
			UserInfo: dal.UserInfo{},
		})
		return
	} else if err != nil { // other errors
		errHandler.RegisterLoginErr(err)
		return
	}
	// record found
	passwordHashByte := []byte(userCredential.Password)
	passwordByte := []byte(req.Password)
	// check password correctness
	if err := bcrypt.CompareHashAndPassword(passwordHashByte, passwordByte); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		zap.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			Code:     eh.CodeLogicalError,
			Msg:      "incorrect password",
			UserInfo: dal.UserInfo{},
		})
		return
	}
	var userInfo dal.UserInfo
	userInfo, err = userInfo.FindByUserId(userCredential.Id)
	if err != nil {
		zap.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			Code:     eh.CodeLogicalError,
			Msg:      "user info not found",
			UserInfo: dal.UserInfo{},
		})
		return
	}
	// generate JWT
	token, err := middleware.GenToken(userCredential.Id, userInfo.Username)
	if err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	cookie.SetToken(c, token)
	cookie.SetUserId(c, strconv.FormatInt(userCredential.Id, 10))
	c.JSON(http.StatusOK, dto.RegisterLoginResp{
		Code:     eh.CodeOK,
		Msg:      "successfully logged in",
		UserInfo: userInfo,
	})
}

func (u *UserCredential) Logout(c *gin.Context) {
	if _, err := cookie.GetUserId(c); err != nil {
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			Code:     eh.CodeTokenError,
			Msg:      "you're not logged in",
			UserInfo: dal.UserInfo{},
		})
		return
	}
	cookie.ClearAllUserInfos(c)
	c.JSON(http.StatusOK, dto.RegisterLoginResp{
		Code:     eh.CodeOK,
		Msg:      "successfully logged out",
		UserInfo: dal.UserInfo{},
	})
}

func (u *UserCredential) UpdatePassword(c *gin.Context) {
	var req dto.RegisterLoginReq
	errHandler := eh.User{C: c}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	userIdStr, err := cookie.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			Code:     eh.CodeTokenError,
			Msg:      "you're not logged in",
			UserInfo: dal.UserInfo{},
		})
		return
	}
	userId := util.ParseId(userIdStr)
	var userInfo dal.UserInfo
	userInfo, err = userInfo.FindByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			Code:     eh.CodeDatabaseError,
			Msg:      "user not exits",
			UserInfo: dal.UserInfo{},
		})
		return
	}
	var userCredential dal.UserCredential
	if err := userCredential.UpdatePassword(userInfo.UserId, req.Password); err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	c.JSON(http.StatusOK, dto.RegisterLoginResp{
		Code:     eh.CodeOK,
		Msg:      "password changed successfully",
		UserInfo: dal.UserInfo{},
	})
}

func bCryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

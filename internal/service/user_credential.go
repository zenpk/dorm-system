package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/middleware"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/eh"
	"github.com/zenpk/dorm-system/pkg/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserCredential struct{}

func (u *UserCredential) Register(c *gin.Context) {
	var req dto.RegisterLoginReq
	errHandler := eh.JSONHandler{C: c, V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.Handle(err)
		return
	}
	// check password length
	if len(req.Password) < viper.GetInt("auth.password_length") {
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeInputError,
			Msg:  "password too short",
		})
		return
	}
	// username duplication check
	var table *dal.UserCredential
	_, err := table.FindByUsername(req.Username)
	if err == nil { // record found
		zap.Logger.Warn("user already exists")
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeDatabaseError,
			Msg:  "user already exists",
		})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) { // not "Not Found" error
		errHandler.Handle(err)
		return
	}
	// no duplication
	passwordHash, err := bCryptPassword(req.Password)
	if err != nil {
		errHandler.Handle(err)
		return
	}
	newUserCredential, _, err := table.RegisterNewUser(req.Username, passwordHash)
	if err != nil {
		errHandler.Handle(err)
		return
	}
	// generate JWT
	token, err := middleware.GenToken(newUserCredential.Id, req.Username)
	if err != nil {
		errHandler.Handle(err)
		return
	}
	cookie.SetToken(c, token)
	cookie.SetUserId(c, strconv.FormatUint(newUserCredential.Id, 10))
	cookie.SetUsername(c, req.Username)
	c.JSON(http.StatusOK, dto.CommonResp{
		Code: eh.Preset.CodeOK,
		Msg:  "successfully registered",
	})
}

func (u *UserCredential) Login(c *gin.Context) {
	var req dto.RegisterLoginReq
	errHandler := eh.JSONHandler{C: c, V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.Handle(err)
		return
	}
	var userCredential *dal.UserCredential
	var err error
	userCredential, err = userCredential.FindByUsername(req.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) { // no record
		zap.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeDatabaseError,
			Msg:  "user not exits",
		})
		return
	} else if err != nil { // other errors
		errHandler.Handle(err)
		return
	}
	// record found
	passwordHashByte := []byte(userCredential.Password)
	passwordByte := []byte(req.Password)
	// check password correctness
	if err := bcrypt.CompareHashAndPassword(passwordHashByte, passwordByte); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		zap.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeLogicalError,
			Msg:  "incorrect password",
		})
		return
	}
	var userInfo *dal.UserInfo
	userInfo, err = userInfo.FindByUserId(userCredential.Id)
	if err != nil {
		zap.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeLogicalError,
			Msg:  "user info not found",
		})
		return
	}
	// generate JWT
	token, err := middleware.GenToken(userCredential.Id, userInfo.Username)
	if err != nil {
		errHandler.Handle(err)
		return
	}
	cookie.SetToken(c, token)
	cookie.SetUserId(c, strconv.FormatUint(userCredential.Id, 10))
	c.JSON(http.StatusOK, dto.CommonResp{
		Code: eh.Preset.CodeOK,
		Msg:  "successfully logged in",
	})
}

func (u *UserCredential) Logout(c *gin.Context) {
	if _, err := cookie.GetUserId(c); err != nil {
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeTokenError,
			Msg:  "you're not logged in",
		})
		return
	}
	cookie.ClearAllUserInfos(c)
	c.JSON(http.StatusOK, dto.CommonResp{
		Code: eh.Preset.CodeOK,
		Msg:  "successfully logged out",
	})
}

func (u *UserCredential) UpdatePassword(c *gin.Context) {
	var req dto.RegisterLoginReq
	errHandler := eh.JSONHandler{C: c, V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.Handle(err)
		return
	}
	if len(req.Password) < viper.GetInt("auth.password_length") {
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeTokenError,
			Msg:  "password too short",
		})
		return
	}
	userIdStr, err := cookie.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeTokenError,
			Msg:  "you're not logged in",
		})
		return
	}
	userId := util.ParseU64(userIdStr)
	var userInfo *dal.UserInfo
	userInfo, err = userInfo.FindByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeDatabaseError,
			Msg:  "user not exits",
		})
		return
	}
	var userCredential *dal.UserCredential
	userCredential, err = userCredential.FindById(userInfo.UserId)
	if err != nil {
		errHandler.Handle(err)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userCredential.Password), []byte(req.Password)); !errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		c.JSON(http.StatusOK, dto.CommonResp{
			Code: eh.Preset.CodeDatabaseError,
			Msg:  "new password cannot be the same as old one",
		})
		return
	}
	// update password
	userCredential.Password = req.Password
	if err := userCredential.Update(userCredential); err != nil {
		errHandler.Handle(err)
		return
	}
	c.JSON(http.StatusOK, dto.CommonResp{
		Code: eh.Preset.CodeOK,
		Msg:  "password changed successfully",
	})
}

func bCryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

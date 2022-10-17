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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"strconv"
)

type UserCredential struct{}

// Register 用户注册
func (u *UserCredential) Register(c *gin.Context) {
	// 自动将表格信息映射到 dto 中
	var req dto.RegisterReq
	errHandler := eh.User{C: c}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	// 检查密码长度
	if len(req.Password) < viper.GetInt("auth.password_length") {
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			SuccessCode: 0,
			Status:      -1,
			Message:     "密码长度过短",
			Data:        dal.UserInfo{},
		})
		return
	}
	// 检查邮箱格式
	matched, err := regexp.Match(`[a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+`, []byte(req.Email))
	if !matched {
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			SuccessCode: 0,
			Status:      -1,
			Message:     "邮箱格式错误",
			Data:        dal.UserInfo{},
		})
		return
	}
	if err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	// 检查邮箱重复
	var userCredentialDal dal.UserCredential
	_, err = userCredentialDal.FindByEmail(req.Email)
	if err == nil { // 已经存在该用户名对应的记录
		middleware.Logger.Warn("user already exists")
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			SuccessCode: 0,
			Status:      2,
			Message:     "用户已存在",
			Data:        dal.UserInfo{},
		})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) { // 出错了，且不是 “未找到用户” 错误
		errHandler.RegisterLoginErr(err)
		return
	}
	// 没有重复，则注册此新用户
	passwordHash, err := bCryptPassword(req.Password) // 密码 BCrypt 加密
	if err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	newUserCredential, newUserInfo, err := userCredentialDal.RegisterNewUser(req.Email, req.Username, passwordHash)
	if err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	// 生成 JWT
	token, err := middleware.GenToken(newUserCredential.Id, req.Username)
	if err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	// 写入 Cookie
	cookie.SetToken(c, token)
	cookie.SetUserId(c, strconv.FormatInt(newUserCredential.Id, 10))
	cookie.SetUsername(c, req.Username)
	c.JSON(http.StatusOK, dto.RegisterLoginResp{
		SuccessCode: 1,
		Status:      0,
		Message:     "注册成功",
		Data:        newUserInfo,
	})
}

// Login 用户登录
func (u *UserCredential) Login(c *gin.Context) {
	var req dto.LoginReq
	errHandler := eh.User{C: c}
	if err := c.ShouldBindJSON(&req); err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	var userCredentialDal dal.UserCredential
	userCredential, err := userCredentialDal.FindByEmail(req.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) { // 没有查到该邮箱对应的记录
		middleware.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			SuccessCode: 0,
			Status:      1,
			Message:     "用户不存在",
			Data:        dal.UserInfo{},
		})
		return
	} else if err != nil { // 出现其他错误
		errHandler.RegisterLoginErr(err)
		return
	}
	// 有对应的用户名，则比较密码
	passwordHashByte := []byte(userCredential.Password)
	passwordByte := []byte(req.Password)
	// 检查密码是否正确
	if err := bcrypt.CompareHashAndPassword(passwordHashByte, passwordByte); errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		middleware.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			SuccessCode: 0,
			Status:      2,
			Message:     "密码错误",
			Data:        dal.UserInfo{},
		})
		return
	}
	userInfo, err := userInfoDal.FindByUserId(userCredential.Id)
	if err != nil {
		middleware.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.RegisterLoginResp{
			SuccessCode: 0,
			Status:      -1,
			Message:     "未找到用户对应信息",
			Data:        dal.UserInfo{},
		})
		return
	}
	// 生成 JWT
	token, err := middleware.GenToken(userCredential.Id, userInfo.Username)
	if err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	// 写入 Cookie
	cookie.SetToken(c, token)
	cookie.SetUserId(c, strconv.FormatInt(userCredential.Id, 10))
	cookie.SetUsername(c, userInfo.Username)
	c.JSON(http.StatusOK, dto.RegisterLoginResp{
		SuccessCode: 1,
		Status:      0,
		Message:     "登录成功",
		Data:        userInfo,
	})
}

// Logout 用户登出
func (u *UserCredential) Logout(c *gin.Context) {
	errHandler := eh.User{C: c}
	if _, err := cookie.GetUserId(c); err != nil {
		errHandler.RegisterLoginErr(err)
		return
	}
	cookie.ClearAllUserInfos(c)
	c.JSON(http.StatusOK, dto.RegisterLoginResp{
		SuccessCode: 1,
		Status:      0,
		Message:     "登出成功",
		Data:        dal.UserInfo{},
	})
}

// LoginState 返回用户的登录状态 // TODO
func (u *UserCredential) LoginState(c *gin.Context) {
	token, err := cookie.GetToken(c)
	if err != nil {
		middleware.Logger.Warn(err.Error())
		c.JSON(http.StatusOK, dto.LoginStateResp{CurrentState: 0})
		return
	}
	if token == "" {
		c.JSON(http.StatusOK, dto.LoginStateResp{CurrentState: 0})
		return
	}
	c.JSON(http.StatusOK, dto.LoginStateResp{CurrentState: 1})
}

// bCryptPassword 密码 BCrypt 加密
func bCryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

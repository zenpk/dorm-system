package cookie

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/jwt"
	"strconv"
)

func GetAccessToken(c *gin.Context) string {
	accessToken, _ := c.Cookie("access_token")
	return accessToken
}

func SetAccessToken(c *gin.Context, token string, urls ...string) {
	url := getURL(urls...)
	setSameSite(c)
	c.SetCookie("access_token", token, viper.GetInt("cookie.access_token_age"), "/", url, viper.GetBool("cookie.secure"), true)
}

func GetRefreshToken(c *gin.Context) string {
	refreshToken, _ := c.Cookie("refresh_token")
	return refreshToken
}

func SetRefreshToken(c *gin.Context, token string, urls ...string) {
	url := getURL(urls...)
	setSameSite(c)
	c.SetCookie("refresh_token", token, viper.GetInt("cookie.refresh_token_age"), "/", url, viper.GetBool("cookie.secure"), true)
}
func GetUserId(c *gin.Context) string {
	userId, _ := c.Cookie("_userId")
	return userId
}

func SetUserId(c *gin.Context, userId string, urls ...string) {
	url := getURL(urls...)
	setSameSite(c)
	c.SetCookie("_userId", userId, 0, "/", url, viper.GetBool("cookie.secure"), true)
}

func GetUsername(c *gin.Context) string {
	username, _ := c.Cookie("_username")
	return username
}

func SetUsername(c *gin.Context, username string, urls ...string) {
	url := getURL(urls...)
	setSameSite(c)
	c.SetCookie("_username", username, 0, "/", url, viper.GetBool("cookie.secure"), true)
}

func GetRole(c *gin.Context) string {
	role, _ := c.Cookie("_role")
	return role
}

func SetRole(c *gin.Context, role string, urls ...string) {
	var url string

	setSameSite(c)
	c.SetCookie("_role", role, 0, "/", url, viper.GetBool("cookie.secure"), true)
}

func ClearAllUserInfos(c *gin.Context, urls ...string) {
	url := getURL(urls...)
	c.SetCookie("access_token", "", -1, "/", url, viper.GetBool("cookie.secure"), true)
	c.SetCookie("refresh_token", "", -1, "/", url, viper.GetBool("cookie.secure"), true)
	c.SetCookie("_userId", "", -1, "/", url, viper.GetBool("cookie.secure"), true)
	c.SetCookie("_username", "", -1, "/", url, viper.GetBool("cookie.secure"), true)
	c.SetCookie("_role", "", -1, "/", url, viper.GetBool("cookie.secure"), true)
}

func SetAllFromAccessToken(c *gin.Context, token string, urls ...string) error {
	claims, err := jwt.ParseToken(token)
	if err != nil {
		return err
	}
	if claims.UserId <= 0 {
		return errors.New("userId cannot be 0")
	}
	SetUserId(c, strconv.FormatUint(claims.UserId, 10))
	SetUsername(c, claims.Username)
	SetRole(c, strconv.FormatInt(int64(claims.Role), 10))
	return nil
}

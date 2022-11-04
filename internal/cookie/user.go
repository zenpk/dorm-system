package cookie

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/jwt"
)

func GetToken(c *gin.Context) (string, error) {
	return c.Cookie("token")
}

func SetToken(c *gin.Context, token string, urls ...string) {
	var url string
	if len(urls) <= 0 {
		url = ""
	} else {
		url = viper.GetString("server.domain") + urls[0]
	}
	c.SetCookie("token", token, viper.GetInt("cookie.token_age"), "/", url, false, true)
}
func GetUserId(c *gin.Context) (string, error) {
	return c.Cookie("_userId")
}

func SetUserId(c *gin.Context, userId string, urls ...string) {
	var url string
	if len(urls) <= 0 {
		url = ""
	} else {
		url = viper.GetString("server.domain") + urls[0]
	}
	c.SetCookie("_userId", userId, 0, "/", url, false, true)
}

func GetUsername(c *gin.Context) (string, error) {
	return c.Cookie("_username")
}

func SetUsername(c *gin.Context, username string, urls ...string) {
	var url string
	if len(urls) <= 0 {
		url = ""
	} else {
		url = viper.GetString("server.domain") + urls[0]
	}
	c.SetCookie("_username", username, 0, "/", url, false, true)
}

func GetRole(c *gin.Context) (string, error) {
	return c.Cookie("_role")
}

func SetRole(c *gin.Context, role string, urls ...string) {
	var url string
	if len(urls) <= 0 {
		url = ""
	} else {
		url = viper.GetString("server.domain") + urls[0]
	}
	c.SetCookie("_role", role, 0, "/", url, false, true)
}

func ClearAllUserInfos(c *gin.Context, urls ...string) {
	var url string
	if len(urls) <= 0 {
		url = ""
	} else {
		url = viper.GetString("server.domain") + urls[0]
	}
	c.SetCookie("token", "", -1, "/", url, false, true)
	c.SetCookie("_userId", "", -1, "/", url, false, true)
	c.SetCookie("_username", "", -1, "/", url, false, true)
}

func SetAllFromToken(c *gin.Context, token string, urls ...string) error {
	var url string
	if len(urls) <= 0 {
		url = ""
	} else {
		url = urls[0]
	}
	userId, username, err := jwt.ParseToken(token)
	if err != nil {
		return err
	} else if userId == "0" {
		return errors.New("user_id cannot be 0")
	}
	SetUserId(c, userId, url)
	SetUsername(c, username, url)
	return nil
}

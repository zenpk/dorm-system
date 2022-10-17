package cookie

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetToken(c *gin.Context) (string, error) {
	return c.Cookie("token")
}

func SetToken(c *gin.Context, token string) {
	c.SetCookie("token", token, viper.GetInt("cookie.token_age"), "/", viper.GetString("server.domain"), false, true)
}
func GetUserId(c *gin.Context) (string, error) {
	return c.Cookie("_userId")
}

func SetUserId(c *gin.Context, userId string) {
	c.SetCookie("_userId", userId, 0, "/", viper.GetString("server.domain"), false, true)
}

func GetUsername(c *gin.Context) (string, error) {
	return c.Cookie("_username")
}

func SetUsername(c *gin.Context, username string) {
	c.SetCookie("_username", username, 0, "/", viper.GetString("server.domain"), false, true)
}

func GetRole(c *gin.Context) (string, error) {
	return c.Cookie("_role")
}

func SetRole(c *gin.Context, role string) {
	c.SetCookie("_role", role, 0, "/", viper.GetString("server.domain"), false, true)
}

// ClearAllUserInfos 清除所有用户相关 Cookie
func ClearAllUserInfos(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", viper.GetString("server.domain"), false, true)
	c.SetCookie("_userId", "", -1, "/", viper.GetString("server.domain"), false, true)
	c.SetCookie("_username", "", -1, "/", viper.GetString("server.domain"), false, true)
}

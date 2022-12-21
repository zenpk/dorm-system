package cookie

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func getURL(urls ...string) string {
	if len(urls) <= 0 {
		return ""
	} else {
		return viper.GetString("server.domain") + urls[0]
	}
}

func setSameSite(c *gin.Context) {
	switch viper.GetString("cookie.same_site") {
	case "none":
		c.SetSameSite(http.SameSiteNoneMode)
	case "lax":
		c.SetSameSite(http.SameSiteLaxMode)
	case "strict":
		c.SetSameSite(http.SameSiteStrictMode)
	default:
		c.SetSameSite(http.SameSiteDefaultMode)
	}
}

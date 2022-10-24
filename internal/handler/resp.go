package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type finalResponse struct {
	v interface{}
}

func response(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
}

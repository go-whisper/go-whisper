package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

type response struct {
	TrackID string      `json:"track_id"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (*Controller) Succeed(c *gin.Context, data interface{}) {
	resp := response{
		Success: true,
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

func (*Controller) Fail(c *gin.Context, data interface{}) {
	resp := response{
		Success: false,
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}

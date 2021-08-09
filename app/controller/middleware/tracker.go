package middleware

import (
	"strings"
	"time"

	"github.com/go-whisper/go-whisper/app/instance"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Tracker() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.RequestURI, ".css") || strings.Contains(c.Request.RequestURI, ".js") {
			return
		}
		t1 := time.Now()
		u1, err := uuid.NewUUID()
		if err != nil {
			instance.Logger().Error("创建UUID失败", zap.String("caller", "middleware.Tracker"), zap.Error(err))
			return
		}
		trackID := u1.String()
		c.Set("track_id", trackID)
		instance.Logger().Info(
			"received request",
			zap.Strings("tags", []string{"middleware.tracker"}),
			zap.String("url", c.Request.Method+":"+c.Request.RequestURI),
			zap.String("from", c.ClientIP()),
			zap.String("track_id", trackID),
		)
		c.Next()
		instance.Logger().Info(
			"request done",
			zap.Strings("tags", []string{"middleware.tracker"}),
			zap.String("track_id", trackID),
			zap.Duration("times", time.Now().Sub(t1)),
		)
	}
}

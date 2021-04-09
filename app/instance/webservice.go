package instance

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine

var onceGinEngine sync.Once

func WebService() *gin.Engine {
	onceGinEngine.Do(func() {
		ginEngine = gin.Default()
	})
	return ginEngine
}

package entrypoints

import (
	"github.com/RafflesApp/imageOptimizer/src/handlers"
	"github.com/gin-gonic/gin"
)

func SetEntryPoints(api *gin.RouterGroup) {
	api.GET("/", handlers.WelcomeApi)
	api.POST("/upload", handlers.Upload)
}

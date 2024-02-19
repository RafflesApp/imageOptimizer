package handlers

import (
	"github.com/RafflesApp/imageOptimizer/src/constants"
	"github.com/gin-gonic/gin"
	"time"
)

type HomeResponse struct {
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

func WelcomeApi(context *gin.Context) {
	context.JSON(200, HomeResponse{Time: time.Now(), Message: constants.WelcomeMessage})
}

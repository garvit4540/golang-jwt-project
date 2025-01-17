package routes

import (
	"github.com/gin-gonic/gin"
)

func AuthRoutes(ctx *gin.Engine) {
	ctx.POST("users/signup", controllers.Signup())
	ctx.POST("users/login", controllers.Login())
}

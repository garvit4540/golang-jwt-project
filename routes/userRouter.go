package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(ctx *gin.Engine) {
	ctx.POST("users/users", controllers.GetUsers())
	ctx.POST("users/:user_id", controllers.GetUser())
}

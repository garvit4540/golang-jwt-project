package routes

import (
	"github.com/garvit4540/golang-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(ctx *gin.Engine) {
	ctx.POST("users/users", controllers.GetUsers())
	ctx.POST("users/:user_id", controllers.GetUser())
}

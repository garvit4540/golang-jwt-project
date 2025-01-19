package routes

import (
	"github.com/garvit4540/golang-jwt-project/controllers"
	"github.com/garvit4540/golang-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(ctx *gin.Engine) {
	ctx.Use(middleware.Authenticate())
	ctx.GET("users", controllers.GetUsers())
	ctx.GET("users/:user_id", controllers.GetUser())
}

package routes

import (
	"assesment/auth"
	"assesment/config"
	"assesment/handler"
	"assesment/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connection()
	userRepository          = user.NewUserRepository(DB)
	userService             = user.NewUserService(userRepository, authService)
	authService             = auth.NewAuthService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRouter(r *gin.Engine) {
	r.GET("/users")
	r.GET("users/:user_id")
	r.POST("users/register")
	r.POST("users/login")
	r.PUT("users/:user_id")
	r.DELETE("users/:user_id")

}

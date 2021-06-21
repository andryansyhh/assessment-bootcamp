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
	r.GET("/users", handler.Middleware(userService, authService), userHandler.GetUserByID)
	r.GET("users/:user_id", handler.Middleware(userService, authService), userHandler.GetUserByID)
	r.POST("users/register", userHandler.UserRegister)
	r.POST("users/login", userHandler.UserLogin)
	r.PUT("users/:user_id", userHandler.UpdateUser)
	r.DELETE("users/:user_id", userHandler.DeleteUser)

}

package routes

import (
	"assesment/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB = config.Connection()

func UserRouter(r *gin.Engine) {
	r.GET("/users")
	r.GET("users/:user_id")
	r.POST("users/register")
	r.POST("users/login")
	r.PUT("users/:user_id")
	r.DELETE("users/:user_id")

}

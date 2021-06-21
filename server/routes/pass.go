package routes

import "github.com/gin-gonic/gin"

func PassRouter(r *gin.Engine) {
	r.GET("/password")
	r.GET("password/:pass_id")
	r.POST("password")
	r.PUT("password/:pass_id")
	r.DELETE("password/:pass_id")
}

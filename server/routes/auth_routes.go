package routes

import (
	"kopikasir-backend/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup){
	authGroup := r.Group("/auth") // untuk masuk routes ini harus diawali /auth
	{
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/login",auth.Login)
	}
}
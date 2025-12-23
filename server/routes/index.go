package routes

import (
	"kopikasir-backend/middleware"

	"github.com/gin-gonic/gin"
)

// setup router sebagai gerbang utama
func SetupRouter(r *gin.Engine){
	
	// membuat grup /api untuk memberitahu ini api
	api := r.Group("/api")
	{
		// route tanpa token
		AuthRoutes(api)

		// route dengan token
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			//
			protected.GET("/profile-test", func(c *gin.Context) {
				// Coba ambil data yang dititip middleware tadi
				userID, _ := c.Get("userID")
				role, _ := c.Get("role")
				
				c.JSON(200, gin.H{
					"message": "Anda berhasil masuk area rahasia!",
					"user_id": userID,
					"role": role,
				}) 
			})
		}
	}
}
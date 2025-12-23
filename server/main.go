package main

import (
	"kopikasir-backend/config"
	"kopikasir-backend/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	
	// Memanggil Koneksi Database
	config.ConnectDatabase()

	r := gin.Default()

	// Testing Route Sederhana
	r.GET("/ping",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"message":"pong",
			"status":"Project KopiKasir Ready!",
		})
	})

	// panggil manager route
	routes.SetupRouter(r)

	// Menjalankan di server pada port 8080
	r.Run(":8080")
}
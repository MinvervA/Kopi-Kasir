package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){

		// 1. mengambil header authorization
		authHeader := c.GetHeader("Authorization")

		// 2. Validasi Format: mewajibkan ada "Bearer <Token>"
		if authHeader == "" || !strings.HasPrefix(authHeader,"Bearer "){
			c.JSON(http.StatusUnauthorized,gin.H{
				"error" : true,
				"message" : "Unauthorized : Token tidak ditemukan!",
			})
			c.Abort()
			return
		}

		// 3. Ambil String Tokennya saja (buang kata "Bearer" )
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 4. Parse & Validasi Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{},error) {
			// pastikan metode enkripsinya sesuai (HMAC)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metode signing tidak valid")
			}
			// kembalikan secret key untuk check key nya
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// jika token rusak atau expired
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized,gin.H{
				"error" : true,
				"message" : "Unauthorized: Token tidak valid !",
			})
			c.Abort()
			return 
		}

		// 5. claim data dari token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// simpan ke context gin ( penyimpanan sementara )
			c.Set("userID", claims["sub"])
			c.Set("role", claims["role"])
		}

		c.Next()
	}
}
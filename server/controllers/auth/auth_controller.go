package auth

import (
	"fmt"
	"kopikasir-backend/config"
	"kopikasir-backend/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username	string		`json:"username" binding:"required"`
	Password	string		`json:"password" binding:"required"`
	Role		string		`json:"role"`
}

type LoginInput struct {
	Username	string		`json:"username" binding:"required"`
	Password	string		`json:"password" binding:"required"`
}

func Register(c *gin.Context){
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// password hasing
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.DefaultCost)

	// membuat struct user
	user := models.User{
			Username:	input.Username,
			Password:	string(hashedPassword),
			Role:		input.Role,
	}

	// simpan ke database
	// logic: INSERT INTO users
	if err := config.DB.Create(&user).Error; err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : true,
			"message" : "Gagal membuat user, username mungkin sudah ada",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"error" : false,
		"message" : "Register berhasil!",
		"data" : user,
	})
}

func Login(c *gin.Context){
	var input LoginInput

	// validasi input pada struct LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : true,
			"error_detail" : err.Error(),
		})
		return
	}
	fmt.Println(input.Username)
	fmt.Println(input.Password)

	// melakukan pencarian user di Database berdasarkan username
	// SQL : SELECT * FROM users WHERE username = "..."
	var user models.User
	if err := config.DB.Where("username = ?",input.Username).First(&user).Error; err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":true,
			"message":"Username tidak di temukan!",
		})
		return
	}

	// check password ( membandingkan yang ada pada db dan input dengan bcrypt)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password)); err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : true,
			"message" : "Password Salah !",
		})
		return
	}

	// mengenerate JWT Token jika username dan password benar
	// settingan token disini 24 jam
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub" : user.ID,
		"role" : user.Role,
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	})

	// melakukan validasi secret key pada env
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : true,
			"message" : "JWT_SECRET Secret key belum di set",
		})
		return
	}

	// ditambahkan dengan secretkey untuk jwtnya
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error" : true,
			"message" : "Gagal Generate Token",
			"error_detail" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"error" : false,
		"message" : "Berhasil Login !",
		"token" : tokenString,
		"role" : user.Role,
	})


}
package controllers

import (
	"main/db"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context){
	// mendapatkan email/password di req body
	var Body struct{
		Nama		string 	`json:"nama"`
		Email 		string
		Password 	string
		Telepon		int		`json:"telepon"`
	}

	if c.ShouldBind(&Body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal mendapatkan request body",
		})

		return
	}

	//hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(Body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal hash password",
		})

		return
	}

	//buat user baru
	user := db.User{Email: Body.Email, Password: string(hash), Nama: Body.Nama, Telepon: Body.Telepon}
	result := db.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal membuat user",
		})

		return
	}

	//reponse
	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil membuat user",
	})
}

func Login(c *gin.Context){
	// mendapatkan email/password di req body
	var Body struct{
		Email 		string
		Password 	string
	}

	if c.Bind(&Body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "gagal mendapatkan request body",
		})

		return
	}

	// melihat request user
	user := db.User{}
	db.DB.First(&user, "email = ?", Body.Email)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email dan password tidak ada",
		})

		return
	}

	// compare password user dengan password hash
	err  := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email dan password salah",
		})

		return
	}

	// generate token jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gagal mendapatkan token",
		})

		return
	}

	//pesan kembali
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil login",
	})
}

func Validasi(c *gin.Context){
	// validasi login
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

package main

import (
	"main/controllers"
	"main/db"
	"main/intializers"
	"main/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	//load env
	initializers.LoadEnv()
	//connect database
    db.ConnectDatabase()

	//init router
	r := gin.Default()

  	r.POST("/register", controllers.Register) // Request register
	r.POST("/login", controllers.Login) // Request login
	r.GET("/validasi",middlewares.Requireauth, controllers.Validasi) // Request validasi login

  	r.Run()
}
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
	r.GET("/me", middlewares.Requireauth, controllers.Validasi) // Request validasi login
	r.GET("/api/product", middlewares.Requireauth, controllers.Index)
	r.GET("/api/product:id", middlewares.Requireauth, controllers.Show)
	r.POST("/api/product", middlewares.Requireauth, controllers.Create)
	r.PUT("/api/product:id", middlewares.Requireauth, controllers.Update)
	r.DELETE("/api/product", middlewares.Requireauth, controllers.Delete)

  	r.Run()
}
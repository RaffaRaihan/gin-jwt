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

	//user
  	r.POST("/register", controllers.Register) // Request register
	r.POST("/login", controllers.Login) // Request login
	r.GET("/me", middlewares.Requireauth, controllers.Validasi) // Request validasi login

	//Crud Role
	r.GET("/api/role", middlewares.Requireauth, controllers.GetRoles)
	r.GET("/api/role/:ID", middlewares.Requireauth, controllers.ReadRoles)
	r.POST("/api/role", middlewares.Requireauth, controllers.CreateRoles)
	r.PUT("/api/role/:ID", middlewares.Requireauth, controllers.UpdateRoles)
	r.DELETE("/api/role", middlewares.Requireauth, controllers.DeleteRoles)

	//Crud Product
	r.GET("/api/product", middlewares.Requireauth, controllers.Index)
	r.GET("/api/product/:ID", middlewares.Requireauth, controllers.Show)
	r.POST("/api/product", middlewares.Requireauth, controllers.Create)
	r.PUT("/api/product/:ID", middlewares.Requireauth, controllers.Update)
	r.DELETE("/api/product", middlewares.Requireauth, controllers.Delete)

	//Crud Ac
	r.GET("/api/ac", middlewares.Requireauth, controllers.GetAc)
	r.GET("/api/ac/:ID", middlewares.Requireauth, controllers.ReadAc)
	r.POST("/api/ac", middlewares.Requireauth, controllers.CreateAc)
	r.PUT("/api/ac/:ID", middlewares.Requireauth, controllers.UpdateAc)
	r.DELETE("/api/ac", middlewares.Requireauth, controllers.DeleteAc)

	//Crud Service
	r.GET("/api/service", middlewares.Requireauth, controllers.GetService)
	r.GET("/api/service/:ID", middlewares.Requireauth, controllers.ReadService)
	r.POST("/api/service", middlewares.Requireauth, controllers.CreateService)
	r.PUT("/api/service/:ID", middlewares.Requireauth, controllers.UpdateService)
	r.DELETE("/api/service", middlewares.Requireauth, controllers.DeleteService)

  	r.Run()
}
package main

import (
	"github.com/Mohamed-Abbas-Homani/jwt-go/controllers"
	"github.com/Mohamed-Abbas-Homani/jwt-go/initializers"
	"github.com/Mohamed-Abbas-Homani/jwt-go/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
	initializers.MigrateDB()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/example", middlewares.Auth, controllers.Example)
	r.Run()
}
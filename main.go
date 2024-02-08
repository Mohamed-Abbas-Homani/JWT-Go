package main

import (
	// "fmt"

	"github.com/Mohamed-Abbas-Homani/jwt-go/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
}

func main() {
	// fmt.Println("Peace")
	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{"message":"Peace"})
	})
	r.Run()
}
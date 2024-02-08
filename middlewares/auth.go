package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Mohamed-Abbas-Homani/jwt-go/initializers"
	"github.com/Mohamed-Abbas-Homani/jwt-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *gin.Context) {
	//Get the cookie from req
	fmt.Println("0")
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	fmt.Println("1")
	//Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2")
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
	//Check the exp
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	fmt.Println("3")
	var user models.User
	result := initializers.DB.First(&user, claims["sub"])
	if result.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Attach to the request
	c.Set("user", user)
	//Continue
	c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
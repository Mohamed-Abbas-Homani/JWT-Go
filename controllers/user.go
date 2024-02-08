package controllers

import (
	"net/http"

	"github.com/Mohamed-Abbas-Homani/jwt-go/initializers"
	"github.com/Mohamed-Abbas-Homani/jwt-go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	//Getting email/password from req body
	var body struct {
		Email 		string
		Password 	string
	}

	if c.Bind(&body) != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error":"Failed to read body."},
		)
		
		return
	}
	//Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error":"Failed to hash password."},
		)
		
		return
	}
	//Create User
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error":"Failed to create user"},
		)
		
		return
	}
	//Response
	c.IndentedJSON(
		http.StatusCreated,
		gin.H{},
	)
}
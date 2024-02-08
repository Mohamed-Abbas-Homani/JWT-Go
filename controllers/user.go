package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Mohamed-Abbas-Homani/jwt-go/initializers"
	"github.com/Mohamed-Abbas-Homani/jwt-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

func Login(c *gin.Context) {
	// Getting Email/Pass from req body
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
	// Look up requested user
	var user models.User
	result := initializers.DB.First(&user, "email = ?", body.Email)
	if result.RowsAffected == 0 {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error":"Invalid Email or Password."},
		)
		
		return
	}
	// Compare send in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error":"Invalid Email or Password."},
		)
		
		return
	}
	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET"))) 
	if err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			gin.H{"error":"Failed to create Token."},
		)
		
		return
	}

	// Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.IndentedJSON(
		http.StatusOK,
		gin.H{},
	)
}


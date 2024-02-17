package handlers

import (
	"jk/go-sportsapp/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Login struct {
	PhoneNumber int    `json:"phone_number"`
	Password    string `json:"password"`
}

func LoginUserHandler(c *gin.Context) {
	c.Request.Header.Add("Access-Control-Allow-Origin", "https://applea.onrender.com")
	i := Login{}
	u := database.User{}

	if err := c.ShouldBindJSON(&i); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Verify number is in correct format
	ok := VerifyNumberIsInCorrectFormat(i.PhoneNumber)
	if !ok {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "invalid mobile number format.",
		})
		return
	}

	// Authenticate and validate user
	u, verified := u.AuthenticateUser(i.PhoneNumber, i.Password, database.DB)
	if !verified {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "mobile number or password do not match.",
		})
		return
	}

	uId, er := u.GetUserID(u.PhoneNumber)
	if er != "" {
		log.Println(er)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "mobile number or password do not match.",
		})
		return
	}

	username := strconv.Itoa(u.PhoneNumber)
	tokenString, err := createJWTToken(username, uId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "authorization failure: " + err.Error(),
		})
		return
	}

	c.Header("Token", tokenString)

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":           uId,
			"phone_number": u.PhoneNumber,
		},
		"token": tokenString,
	})
}

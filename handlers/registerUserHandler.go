package handlers

import (
	"jk/go-sportsapp/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(c *gin.Context) {
	u := database.User{}

	if err := c.ShouldBindJSON(&u); err == nil {
		ok := verifyPasswordMatch(u.Password, u.ConfirmPassword)
		if !ok {
			c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
				"error": "Passwords do not match.",
			})
			return
		}
	}

	// TODO:
	// Hash password before saving to Database
	u, e := u.SaveUser(database.DB)
	if e != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"error": "User is already registered. Please login to continue.",
		})
		return
	}

	username := strconv.Itoa(u.PhoneNumber)

	tokenString, err := createJWTToken(username)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Authorization failure: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"phone_number": u.PhoneNumber,
		},
		"token": tokenString,
	})

}

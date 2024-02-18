package handlers

import (
	"jk/go-sportsapp/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "https://findmyiphone.vercel.app/")
	c.Header("Access-Control-Allow-Headers", "Content-Type,Access-Control-Allow-Origin, Access-Control-Allow-Headers")

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
			"error": e.Error(),
		})
		return
	}

	username := strconv.Itoa(u.PhoneNumber)

	tokenString, err := createJWTToken(username, int(u.Id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Authorization failure: " + err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"user": gin.H{
			"id":           u.Id,
			"phone_number": u.PhoneNumber,
		},
		"token": tokenString,
	})

}

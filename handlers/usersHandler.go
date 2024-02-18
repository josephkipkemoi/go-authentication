package handlers

import (
	"jk/go-sportsapp/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersHandler(c *gin.Context) {
	u := database.User{}
	usrs := u.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": usrs,
	})
}

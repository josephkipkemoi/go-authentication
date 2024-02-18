package handlers

import (
	"jk/go-sportsapp/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersHandler(c *gin.Context) {
	usrs := database.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": usrs,
	})
}

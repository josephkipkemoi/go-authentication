package handlers

import (
	"jk/go-sportsapp/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BalanceHandler(c *gin.Context) {
	usrs := database.User{}
	usrs.GetUsers(database.DB)

	c.JSON(http.StatusOK, gin.H{
		"users": usrs,
	})
}

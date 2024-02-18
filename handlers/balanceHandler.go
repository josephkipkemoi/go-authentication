package handlers

import (
	"jk/go-sportsapp/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BalanceHandler(c *gin.Context) {
	usrs := database.User{}
	u := usrs.GetUsers()

	c.JSON(http.StatusOK, gin.H{
		"users": u,
	})
}

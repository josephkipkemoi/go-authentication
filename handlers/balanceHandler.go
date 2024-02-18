package handlers

import (
	"jk/go-sportsapp/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BalanceHandler(c *gin.Context) {
	u := database.GetUsers()

	c.JSON(http.StatusOK, gin.H{
		"users": u,
	})
}

package handlers

import (
	"jk/go-sportsapp/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "https://findmyiphone.vercel.app")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Token, Accept, X-Requested-With, Access-Control-Allow-Origin, Access-Control-Allow-Headers")

	usrs := database.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": usrs,
	})
}

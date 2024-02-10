package handlers

import (
	"github.com/gin-gonic/gin"
)

func LandingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome: PinaclebetAPI Ver: 2.0",
	})
}

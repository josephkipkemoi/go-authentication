package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateUserHandler(c *gin.Context) {
	t := c.Request.Header.Get("Authorization")
	token := strings.Split(t, "Bearer ")

	err, usr := VerifyToken(token[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"phone_number": usr,
		},
	})
}

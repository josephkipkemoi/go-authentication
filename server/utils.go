package server

import (
	"jk/go-sportsapp/handlers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() func(*gin.Context) {
	return func(c *gin.Context) {
		authHeader := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized/malformed token",
			})
			return
		} else {
			jwtToken := authHeader[1]
			s := strings.Trim(jwtToken, " ")

			_, err := handlers.VerifyToken(s)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "unauthorized: " + err.Error(),
				})
				return
			}
		}
	}
}

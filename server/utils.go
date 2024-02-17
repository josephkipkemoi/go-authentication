package server

import (
	"jk/go-sportsapp/handlers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetHeaders() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json:charset=utf-8")
		ctx.Header("Host", ctx.Request.Host)
		ctx.Header("X-Powered-By", "Golang")
		ctx.Header("Access-Control-Allow-Origin", "https://applea.onrender.com")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Token")
	}
}

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

			err, _ := handlers.VerifyToken(s)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "unauthorized: " + err.Error(),
				})
				return
			}
		}
	}
}

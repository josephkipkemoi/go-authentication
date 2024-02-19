package server

import (
	"jk/go-sportsapp/handlers"

	"github.com/gin-gonic/gin"
)

func ConnectServerRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", handlers.LandingHandler)
	r.POST("api/register", handlers.RegisterUserHandler)
	r.POST("api/login", handlers.LoginUserHandler)
	r.GET("api/users", handlers.UsersHandler)

	// Token MUST be prvided for below API routes
	r.Use(AuthorizeJWT())

	r.GET("api/user", handlers.AuthenticateUserHandler)
	r.GET("api/users/:user_id/balance", handlers.BalanceHandler)

	return r
}

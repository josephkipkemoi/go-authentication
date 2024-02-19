package server

import (
	"jk/go-sportsapp/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConnectServerRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", handlers.LandingHandler)
	r.POST("api/register", handlers.RegisterUserHandler)
	r.POST("api/login", handlers.LoginUserHandler)
	r.GET("api/users", handlers.UsersHandler)

	r.Use(SetHeaders())
	// Token MUST be prvided for below API routes
	// r.Use(AuthorizeJWT())

	r.GET("api/user", handlers.AuthenticateUserHandler)
	r.GET("api/users/:user_id/balance", handlers.BalanceHandler)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin, Content-Type, Token, Accept, X-Requested-With, withCredentials, Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Origin, Content-Type, Token, Accept, X-Requested-With, Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return r
}

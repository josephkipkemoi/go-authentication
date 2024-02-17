package server

import (
	"jk/go-sportsapp/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConnectServerRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://findmyiphone.vercel.app"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(SetHeaders())

	apiRoutes(r) // API endpoint URLs

	return r
}

func apiRoutes(r *gin.Engine) {
	r.GET("/", handlers.LandingHandler)
	r.POST("api/register", handlers.RegisterUserHandler)
	r.POST("api/login", handlers.LoginUserHandler)

	r.Use(AuthorizeJWT())
	// Token MUST be prvided for below API routes

	r.GET("api/user", handlers.AuthenticateUserHandler)
	r.GET("api/users/:user_id/balance", handlers.BalanceHandler)
}

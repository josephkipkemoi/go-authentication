package server

import (
	"jk/go-sportsapp/handlers"

	"github.com/gin-gonic/gin"
)

func ConnectServerRouter() *gin.Engine {
	r := gin.Default()
	r.Use(SetHeaders())

	apiRoutes(r) // API endpoint URIs

	return r
}

func apiRoutes(r *gin.Engine) {

	r.GET("/", handlers.LandingHandler)
	r.POST("api/register", handlers.RegisterUserHandler)
	r.POST("api/login", handlers.LoginUserHandler)

	r.Use(AuthorizeJWT())

	r.GET("api/users/:user_id/balance", handlers.BalanceHandler)

}

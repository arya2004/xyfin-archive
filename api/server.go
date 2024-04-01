package api

import (
	"fmt"

	database "github.com/arya2004/Xyfin/database/sqlc"
	"github.com/arya2004/Xyfin/token"
	"github.com/arya2004/Xyfin/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	configuration util.Configuration
	store database.Store
	tokenCreator token.Creator
	router *gin.Engine
}


//create a new HTTP server and setup routing
func NewServer(configuration util.Configuration , store database.Store) (* Server, error) {
	tokenMaker, err := token.NewPasetoMaker(configuration.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		configuration: configuration,
		store: store,
		tokenCreator: tokenMaker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenCreator))


	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts/", server.listAccount)
	authRoutes.POST("/transfers", server.createTransfer)

	//Add the router to routes
	server.router = router
	return server, nil

}
//start the server
func (server * Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}
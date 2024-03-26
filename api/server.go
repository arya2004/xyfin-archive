package api

import (
	database "github.com/arya2004/Xyfin/database/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store database.Store
	router *gin.Engine
}


//create a new HTTP server and setup routing
func NewServer(store database.Store) * Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/users", server.createUser)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)
	router.POST("/transfers", server.createTransfer)

	//Add the router to routes
	server.router = router
	return server

}
//start the server
func (server * Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{"error": err.Error()}
}
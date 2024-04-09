package grpcapi

import (
	"fmt"

	database "github.com/arya2004/Xyfin/database/sqlc"
	"github.com/arya2004/Xyfin/pb"
	"github.com/arya2004/Xyfin/token"
	"github.com/arya2004/Xyfin/util"
)



type Server struct {
	pb.UnimplementedXyfinServer
	configuration util.Configuration
	store database.Store
	tokenCreator token.Creator
	
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

	return server, nil

}

package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/arya2004/Xyfin/api"
	database "github.com/arya2004/Xyfin/database/sqlc"
	"github.com/arya2004/Xyfin/grpcapi"
	"github.com/arya2004/Xyfin/pb"
	"github.com/arya2004/Xyfin/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)



func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cant load config")
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := database.NewStore(conn)
	runGRPCServer(config, store)
}

func runGRPCServer(config util.Configuration, store database.Store){
	
	grpcServer := grpc.NewServer()
	server, err := grpcapi.NewServer(config,store)

	if err != nil{
		log.Fatal("cant create server:", err)
	}
	pb.RegisterXyfinServer(grpcServer, server)
	reflection.Register(grpcServer);

	listner, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listner")
	}

	log.Printf("started gRPC @%s", listner.Addr().String())
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal("cant start server")
	}

	
}


func runGinServer(config util.Configuration, store database.Store){
	server, err := api.NewServer(config,store)

	if err != nil{
		log.Fatal("cant create server:", err)
	}

	err = server.Start(config.HTTPSServerAddress)
	if err != nil{
		log.Fatal("Cant start the server:", err)
	}
}
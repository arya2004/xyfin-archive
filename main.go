package main

import (
	"database/sql"
	"log"

	"github.com/arya2004/Xyfin/api"
	database "github.com/arya2004/Xyfin/database/sqlc"
	"github.com/arya2004/Xyfin/util"
	_ "github.com/lib/pq"
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
	server, err := api.NewServer(config,store)

	if err != nil{
		log.Fatal("cant create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil{
		log.Fatal("Cant start the server:", err)
	}
}
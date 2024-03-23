package main

import (
	"database/sql"
	"log"

	"github.com/arya2004/Xyfin/api"
	database "github.com/arya2004/Xyfin/database/sqlc"
	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dbSource = "postgresql://root:secret@localhost:5432/xyfin?sslmode=disable"
const serverAddress = "0.0.0.0:8081"


func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := database.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil{
		log.Fatal("Cant start the server")
	}
}
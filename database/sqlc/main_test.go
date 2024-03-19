package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

//By convention, tests begin with `Test`


const dbDriver = "postgres"
const dbSource = "postgresql://root:secret@localhost:5432/xyfin?sslmode=disable"

var testQueries *Queries



func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = New(conn)

	//Run the unit tests
	os.Exit(m.Run())

}
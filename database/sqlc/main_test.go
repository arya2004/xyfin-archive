package database

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/arya2004/Xyfin/util"
	_ "github.com/lib/pq"
)

//By convention, tests begin with `Test`



var testQueries *Queries
var testDb *sql.DB


func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load config file")
	}

	testDb, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = New(testDb)

	//Run the unit tests
	os.Exit(m.Run())

}
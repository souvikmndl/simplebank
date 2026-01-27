package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/souvikmndl/simplebank/api"
	db "github.com/souvikmndl/simplebank/db/sqlc"
	"github.com/souvikmndl/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("unable to read config %+v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("error starting server")
	}
}

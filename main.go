package main

import (
	"database/sql"
	"log"

	"github.com/gerardoavelar/simplebank/api"
	db "github.com/gerardoavelar/simplebank/db/sqlc"
	"github.com/gerardoavelar/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server ", err)
	}

}

package main

import (
	"database/sql"
	"log"

	"github.com/beso768/Go-BankTansactions/api"
	db "github.com/beso768/Go-BankTansactions/db/sqlc"
	"github.com/beso768/Go-BankTansactions/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(&store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

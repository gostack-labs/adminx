package main

import (
	"database/sql"
	"log"

	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/api"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("connot load config:", err)
	}
	conn, err := sql.Open(config.DB.Driver, config.DB.Source)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.Server.Addr)
	if err != nil {
		log.Fatal("connot start server:", err)
	}
}

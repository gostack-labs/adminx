package main

import (
	"database/sql"
	"log"

	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/api"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/repository/redis"
	_ "github.com/lib/pq"
)

func main() {
	configs.LoadConfig()
	cache, err := redis.New()
	if err != nil {
		log.Fatal("connot connect to redis:", err)
	}

	conn, err := sql.Open(configs.Config.DB.Driver, configs.Config.DB.Source)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(store, cache)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(configs.Config.Server.Addr)
	if err != nil {
		log.Fatal("connot start server:", err)
	}
}

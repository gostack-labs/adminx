package main

import (
	"context"
	"log"

	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/api"
	"github.com/gostack-labs/adminx/internal/middleware/permission"
	db "github.com/gostack-labs/adminx/internal/repository/db/sqlc"
	"github.com/gostack-labs/adminx/internal/repository/redis"
	"github.com/jackc/pgx/v4"
)

//@title adminx
//@service adminx
//@desc adminx服务相关接口
//@baseurl /
func main() {
	configs.LoadConfig()
	cache, err := redis.New()
	if err != nil {
		log.Fatal("connot connect to redis:", err)
	}

	permission.Casbin()

	conn, err := pgx.Connect(context.Background(), configs.Get().DB.Source)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(store, cache)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(configs.Get().Server.Addr)
	if err != nil {
		log.Fatal("connot start server:", err)
	}
}

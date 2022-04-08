package db

import (
	"database/sql"
	"log"

	"github.com/gostack-labs/adminx/configs"
	_ "github.com/lib/pq"
)

type Store interface {
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore() Store {
	db, err := sql.Open(configs.Cfg.DB.Driver, configs.Cfg.DB.Source)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

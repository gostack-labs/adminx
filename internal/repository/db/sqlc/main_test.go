package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/gostack-labs/adminx/configs"
	"github.com/jackc/pgx/v4"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	configs.LoadConfig()
	log.Print(configs.Get().DB.Source)
	testDB, err := pgx.Connect(context.Background(), configs.Get().DB.Source)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

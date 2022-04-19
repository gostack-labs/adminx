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
	conn, err := pgx.Connect(context.Background(), configs.Config.DB.Source)
	if err != nil {
		log.Fatal("connot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}

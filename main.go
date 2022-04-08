package main

import (
	"log"

	"github.com/gostack-labs/adminx/bootstrap"
	"github.com/gostack-labs/adminx/internal/api"
)

func init() {
	bootstrap.Initialize()
}

func main() {
	s, err := api.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = s.Start()
	if err != nil {
		log.Fatal("connot start server:", err)
	}
}

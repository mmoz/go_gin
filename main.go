package main

import (
	"context"
	"log"
	"mmoz/crud/config"
	"mmoz/crud/pkg/database"
	"mmoz/crud/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
		panic(err)
	}

	db, err := database.DbConn(ctx, cfg)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		panic(err)
	}
	s := server.StartServer(db)
	s.Router.Use(config.CORSMiddleware())
	s.StartAuthServer()
	s.StartUserServer()
	s.Router.Run(":4000")
}

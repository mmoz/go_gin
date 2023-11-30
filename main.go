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

	db, err := database.DbConn(ctx)
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

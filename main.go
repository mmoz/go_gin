package main

import (
	"context"
	"log"
	"mmoz/crud/config"
	"mmoz/crud/pkg/database"
	"mmoz/crud/server"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	ctx := context.Background()

	r := gin.Default()
	r.Use(config.CORSMiddleware())
	db, err := database.DbConn(ctx)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		panic(err)
	}
	server.StartAuthServer(r, db)
	server.StartUserServer(r, db)
	r.Run(":4000")
}

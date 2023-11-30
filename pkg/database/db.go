package database

import (
	"context"
	"database/sql"
	"fmt"
	"mmoz/crud/config"
	"time"
)

func DbConn(pctx context.Context, cfg *config.Config) (*sql.DB, error) {

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName))

	if err != nil {
		fmt.Println("Error opening database: ", err)
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println("Error pinging database: ", err)
		db.Close()
		return nil, err
	}

	return db, nil
}

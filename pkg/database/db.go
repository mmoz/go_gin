package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"
)

func DbConn(pctx context.Context) (*sql.DB, error) {

	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	dbUrl := os.Getenv("DB_URL")
	_ = dbUrl

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
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

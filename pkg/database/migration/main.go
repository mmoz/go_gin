package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database connection parameters
	dbUsername := "root"
	dbPassword := "root"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "dbName"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Resolve the absolute path to the migration file
	absPath, err := filepath.Abs("./pkg/database/migration/create_users_table.up.sql")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Absolute Path to Migration:", absPath)

	// Read the content of the SQL file
	sqlContent, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the SQL commands
	_, err = db.Exec(string(sqlContent))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration successful")
}

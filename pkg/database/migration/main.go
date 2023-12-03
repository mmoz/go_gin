package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Database connection parameters
	dbUsername := "root"
	dbPassword := "root1234"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "crud_test"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	// Open a connection to the database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Resolve the absolute path to the migration directory
	migrationDir := "./pkg/database/migration"
	absPath, err := filepath.Abs(migrationDir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Absolute Path to Migration Directory:", absPath)

	// Read all files in the migration directory
	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		log.Fatal(err)
	}

	// Sort files based on their names
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// Execute the SQL commands for each file
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".up.sql") {
			filePath := filepath.Join(absPath, file.Name())
			fmt.Printf("Applying migration from file: %s\n", filePath)

			// Read the content of the SQL file
			sqlContent, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Fatal(err)
			}

			// Execute the SQL commands
			_, err = db.Exec(string(sqlContent))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("Migration successful")
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument.")
		return
	}

	// Load environment variables
	dbUser := getEnv("DB_USER", "test")
	dbPass := getEnv("DB_PASSWORD", "test")
	dbHost := getEnv("DB_HOST", "mysql")
	dbName := getEnv("DB_NAME", "test")

	// Initialize database connection
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		dbUser, dbPass, dbHost, dbName)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	if err := insertData(); err != nil {
		log.Fatalf("Error inserting data: %v", err)
	}
}

func insertData() error {
	arg := os.Args[1] // Get the first command-line argument (index 0 is the program name)
	inserts, convErr := strconv.Atoi(arg)

	if convErr != nil {
		return fmt.Errorf("Error converting '%s' to int: %v\n", arg, convErr)
	}

	// Insert data in batches
	stmt, err := db.Prepare("INSERT INTO test (name) VALUES (?)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %v", err)
	}

	for i := 1; i <= inserts; i++ {
		_, err = stmt.Exec(fmt.Sprintf("test %d", i))

		if i%5_000 == 0 {
			log.Printf("Inserted %d/%d records (%.1f%%)",
				i,
				inserts,
				float64(i)/float64(inserts)*100)
		}
	}
	defer stmt.Close()

	return nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

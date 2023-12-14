package db

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

var db *sql.DB

func Init(){
	connectionString := loadProperties();
	var err error;
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

	initDB()
}

func GetDB() *sql.DB {
	return db
}

func Close() {
    if db != nil {
        db.Close()
    }
}

func loadProperties() string{
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	ip := os.Getenv("DB_IP")
	port := os.Getenv("DB_PORT")

	connectionString := (username+":"+password+"@tcp("+ip+":"+port+")/");
	fmt.Println(connectionString)
	return connectionString;
}

func initDB() {
    data, err := os.Open("./db/database.sql")
    if err != nil {
        log.Fatal(err)
    }
    defer data.Close()

    scanner := bufio.NewScanner(data)
    scanner.Split(bufio.ScanLines)

    var statement string
    for scanner.Scan() {
        line := scanner.Text()
        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "--") {
            continue
        }

        statement += line
        // If line ends with a semicolon, execute the statement
        if strings.HasSuffix(line, ";") {
            _, err := db.Exec(statement)
            if err != nil {
                log.Printf("Error executing statement: %q: %v", statement, err)
            }
            statement = "" // Reset statement
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading from file: %v", err)
    }
    // Switch to the new database using the environment variable
    database := os.Getenv("DATABASE")
    if database != "" {
        _, err := db.Exec("USE " + database)
        if err != nil {
            log.Fatalf("Error selecting database %q: %v", database, err)
        }
    }
    fmt.Println("Database initialized successfully")
}
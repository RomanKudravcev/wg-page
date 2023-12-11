package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error while loading .env")
	}
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	database := os.Getenv("DATABASE")

	fmt.Println("username: ",username)
	connectionString := (username+":"+password+"@tcp("+ip+":"+port+")/"+database);
	return connectionString;
}
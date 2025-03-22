package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var db *sql.DB

func InitDB(){

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),)

	fmt.Println("Waiting for the database Start Up....")
	time.Sleep(5*time.Second)

	db, err := sql.Open("postgres", connStr)
	if err!= nil{
		log.Fatalf("Error Opening database: %v", err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatalf("Error Connecting to the database : %v", err)
	}

	fmt.Println("Successfully Connected to the database")
}

func GetDB() *sql.DB{
	return db
}

func CloseDB(){
	if err := db.Close(); err != nil{
		log.Fatal("Error Closing the Database: %v", err)
	}
}

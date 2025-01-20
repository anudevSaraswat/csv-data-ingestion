package dbconnect

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type connectInfo struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

var db *sql.DB

// this function provides a connection handle to postgresql database server
func ConnectToDatabase() *sql.DB {

	var err error

	if db != nil {
		return db
	}

	connectInfo := connectInfo{
		Host:     os.Getenv("PGSQL_HOST"),
		User:     os.Getenv("PGSQL_USER"),
		Password: os.Getenv("PGSQL_PASSWORD"),
		DBName:   os.Getenv("PGSQL_DB_NAME"),
	}

	connectInfo.Port, err = strconv.Atoi(os.Getenv("PGSQL_PORT"))
	if err != nil {
		// use default
		connectInfo.Port = 5432
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		connectInfo.Host, connectInfo.Port, connectInfo.User, connectInfo.Password, connectInfo.DBName)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln("(ConnectToDatabase) err in sql.Open:", err)
		panic(err)
	}

	return db

}

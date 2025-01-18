package dbconnect

import (
	"database/sql"
	"fmt"
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

// this function provides a connection handle to postgresql database server
// TODO:  convert this to singleton
func ConnectToDatabase() *sql.DB {

	var err error

	connectInfo := connectInfo{
		Host:     os.Getenv("PGSQL_HOST"),
		User:     os.Getenv("PGSQL_USER"),
		Password: os.Getenv("PGSQL_PASSWORD"),
		DBName:   os.Getenv("PGSQL_DB_NAME"),
	}

	connectInfo.Port, err = strconv.Atoi(os.Getenv("PGSQL_PORT"))
	if err != nil {
		panic(err)
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		connectInfo.Host, connectInfo.Port, connectInfo.User, connectInfo.Password, connectInfo.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	return db

}

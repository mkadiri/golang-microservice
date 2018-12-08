package database

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var Db *sql.DB

func init() {
	fmt.Println("Init database")

	var err error
	Db, err = sql.Open("mysql", getDatabaseSourceName())

	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}

	// have a read
	// https://www.alexedwards.net/blog/configuring-
	// https://github.com/go-sql-driver/mysql/issues/461
	Db.SetMaxIdleConns(0)
	Db.SetMaxOpenConns(1)
	Db.SetConnMaxLifetime(time.Minute * 5)
}

func getDatabaseSourceName() string {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	return user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"
}
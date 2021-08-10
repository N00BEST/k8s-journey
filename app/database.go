package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConnection struct {
	Connection              *sql.DB
	isConnectionEstablished bool
}

func getConnectionCreds() string {
	user := os.Getenv("MYSQL_USER")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	log.Println(fmt.Sprintf("Trying to connect to %s@%s:%s", user, host, port))
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user,
		os.Getenv("MYSQL_PASSWORD"),
		host,
		port,
		os.Getenv("MYSQL_DATABASE"))
}

func (db *DatabaseConnection) OpenConnection() error {
	if db.isConnectionEstablished {
		return nil
	}

	var err error
	connectionString := getConnectionCreds()
	db.Connection, err = sql.Open("mysql", connectionString)

	if err != nil {
		db.isConnectionEstablished = false
	} else {
		err = db.Connection.Ping()
		db.isConnectionEstablished = err == nil
	}

	return err
}
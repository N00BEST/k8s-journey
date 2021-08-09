package main

import (
	"database/sql"
	"fmt"
	"os"
)

type DatabaseConnection struct {
	Connection              *sql.DB
	isConnectionEstablished bool
}

func getConnectionCreds() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
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
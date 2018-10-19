package main

import (
	"database/sql"
	"fmt"
)

func connect() (db *sql.DB) {
	connection := getConnection()

	openStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		connection.UserName,
		connection.Password,
		connection.Host,
		connection.Port,
		connection.Database)

	db, err := sql.Open("mysql", openStr)

	checkError(err)

	return db
}
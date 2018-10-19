package main

import (
	"database/sql"
)

func connect() (db *sql.DB) {
	connection := getConnection()

	openStr := connection.UserName +
		":" + connection.Password +
		"@tcp(" + connection.Host +
		":" + connection.Port +
		")/" + connection.Database

	db, err := sql.Open("mysql", openStr)

	checkError(err)

	return db
}
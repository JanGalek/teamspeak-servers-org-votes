package main

import (
	"database/sql"
	"io/ioutil"
	"os"
	"strings"
)

func checkNeedInstall(db *sql.DB) {
	tableExists, err := db.Query("SELECT COUNT(TABLE_NAME) FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_NAME IN('votes', 'voters');")

	checkError(err)

	for tableExists.Next() {
		var count int
		err = tableExists.Scan(&count)

		checkError(err)

		if count != 2 {
			install(db)
		}
	}
}

func install(db *sql.DB)  {

	sqlFile, err := os.Open("./sql/create_tables.sql")
	checkError(err)

	sqlContent, err := ioutil.ReadAll(sqlFile)
	checkError(err)

	requests := strings.Split(string(sqlContent), ";")

	for _, request := range requests {
		if len(strings.TrimSpace(request)) != 0 {
			_, err := db.Exec(request)
			checkError(err)
		}
	}
}
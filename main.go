package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	data := loadData()

	db := connect()
	loadCurrentMonth(db, data)

	updateVoters(db)
	defer db.Close()
}
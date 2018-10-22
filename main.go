package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	data := loadData("current")
	dataPrevious := loadData("previous")
	db := connect()

	checkNeedInstall(db)
	loadCurrentMonth(db, data)
	loadCurrentMonth(db, dataPrevious)

	updateVoters(db)
	defer db.Close()
}
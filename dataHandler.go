package main

import (
	"database/sql"
)

func loadCurrentMonth(db *sql.DB, server Server)  {
	for i := 0; i < len(server.Voters); i++ {
		voter := server.Voters[i]
		insert, insertError := db.Query("INSERT INTO votes (username, votes, month) VALUES (?, ?, ?) ON DUPLICATE KEY UPDATE votes = ?", voter.Nickname, voter.Votes, server.Month, voter.Votes)

		if insertError != nil {
			panic(insertError.Error())
		} else {
			insert.Close()
		}
	}
}

func getSummary(db *sql.DB) (*sql.Rows, error)  {
	summary, err := db.Query("SELECT username, SUM(votes.votes) votes FROM votes GROUP BY username")
	checkError(err)

	return summary, err
}

func updateVoters(db *sql.DB)  {
	summary, err := getSummary(db)

	for summary.Next() {
		var username string
		var votes int

		err = summary.Scan(&username, &votes)

		checkError(err)

		insert, insertError := db.Query("INSERT INTO voters (username, votes) VALUES (?, ?) ON DUPLICATE KEY UPDATE votes = ?", username, votes, votes)

		if insertError != nil {
			panic(insertError.Error())
		} else {
			insert.Close()
		}
	}
}
package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB_URL = os.Getenv("INSTAGRAM_DB")
var DB_PASS = os.Getenv("INSTAGRAM_DB_PASS")

type Statistic struct {
	Date      time.Time
	Follows   int
	Followers int
}

func getStatistics() []Statistic {
	connection_url := fmt.Sprintf("root:%v@tcp(%v:3306)/instagram_statistics?parseTime=true", DB_PASS, DB_URL)

	db, err := sql.Open("mysql", connection_url)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT date, follows, followers FROM Statistics")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	statistics := make([]Statistic, 0)
	var (
		date      time.Time
		follows   int
		followers int
	)
	for rows.Next() {

		rows.Scan(&date, &follows, &followers)
		statistics = append(statistics, Statistic{Date: date, Follows: follows, Followers: followers})
	}

	return statistics
}

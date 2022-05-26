package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//automatic cron job for DB sync
func DBSyncCronJob() {
	db, _ = DBConnection()
	log.Printf("Successfully connected to database")
	_ = CreateTable(db)
	netflix_titles = ReadCSV()
	InsertShows(db, netflix_titles)
}

package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//automatic cron job for DB sync
func DBSyncCronJob() {
	var err error
	db, err = DBConnection()
	if err != nil {
		log.Printf("%s when getting db connection", err)
		return
	}
	log.Printf("Successfully connected to database")

	err = CreateTable(db)
	if err != nil {
		log.Printf("Create netflix_show_details table failed with error %s", err)
		return
	}
	netflix_titles = ReadCSV()
	InsertShows(db, netflix_titles)
}

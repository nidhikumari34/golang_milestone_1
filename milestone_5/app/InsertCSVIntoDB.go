package app

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func InsertCSVIntoDB(w http.ResponseWriter, r *http.Request) {
	netflix_titles = ReadCSV()
	db, _ = DBConnection()
	log.Printf("Successfully connected to database")
	_ = CreateTable(db)
	_ = InsertShows(db, netflix_titles)

}

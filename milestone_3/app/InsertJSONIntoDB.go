package app

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func InsertJSONIntoDB(w http.ResponseWriter, r *http.Request) {
	var err error
	netflix_titles = ReadCSV()

	db, err = DBConnection()
	if err != nil {
		log.Printf("%s when getting db connection", err)
		return
	}
	//defer db.Close()
	log.Printf("Successfully connected to database")

	err = CreateTable(db)
	if err != nil {
		log.Printf("Create netflix_show_details table failed with error %s", err)
		return
	}

	//DBSync()

	err = InsertShowsJSON(r, db, netflix_titles)
	if err != nil {
		log.Printf("Inserting netflix_show_details failed for few or all records")
		return
	}

}

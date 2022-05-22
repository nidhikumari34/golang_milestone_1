package app

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//struct for show details
type Netflix_shows struct {
	Show_id      string `json:"show_id"`
	Show_type    string `json:"show_type"`
	Title        string `json:"title"`
	Director     string `json:"director"`
	Cast         string `json:"cast"`
	Country      string `json:"country"`
	Date_added   string `json:"date_added"`
	Release_year string `json:"release_year"`
	Rating       string `json:"rating"`
	Duration     string `json:"duration"`
	Listed_in    string `json:"listed_in"`
	Description  string `json:"description"`
}

//fetch shows from DB
func ReadDataFromDB() []Netflix_shows {
	var n Netflix_shows
	var tvshows []Netflix_shows

	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("%s when opening DB", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
	}
	log.Printf("Connected to DB %s successfully\n", dbname)

	rows, err := db.Query("select * from netflix_show_details_api")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&n.Show_id, &n.Show_type, &n.Title, &n.Director, &n.Cast, &n.Country, &n.Date_added, &n.Release_year, &n.Rating, &n.Duration, &n.Listed_in, &n.Description)
		if err != nil {
			log.Fatal(err)
		}
		tvshows = append(tvshows, n)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Fetched shows from to DB %s successfully\n", dbname)
	return tvshows
}

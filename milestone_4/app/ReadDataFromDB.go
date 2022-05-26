package app

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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

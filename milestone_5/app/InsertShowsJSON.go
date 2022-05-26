package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//inserting rows in netflix_show_details table
func InsertRowsJSON(db *sql.DB, p Netflix_shows) error {
	query := "INSERT INTO netflix_show_details_api(show_id,show_type,title,director,cast,country,date_added,release_year,rating,duration,listed_in,description) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("%s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx,
		p.Show_id,
		p.Show_type,
		p.Title,
		p.Director,
		p.Cast,
		p.Country,
		p.Date_added,
		p.Release_year,
		p.Rating,
		p.Duration,
		p.Listed_in,
		p.Description)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("%s when finding rows affected", err)
		return err
	}
	log.Printf("%d netflix_show_details created: show_id = %s ", rows, p.Show_id)
	return nil
}

func InsertShowsJSON(w http.ResponseWriter, r *http.Request, db *sql.DB, netflix_titles [][]string) error {
	var err error
	decoder := json.NewDecoder(r.Body)
	var n []Netflix_shows
	err = decoder.Decode(&n)
	if err != nil {
		panic(err)
	}
	for i, _ := range n {
		err = InsertRowsJSON(db, n[i])
		if err != nil {
			log.Printf("Insert failed with %s", err)
		}
	}
	return err
}

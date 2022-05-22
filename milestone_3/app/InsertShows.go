package app

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//inserting rows in netflix_show_details table
func InsertRows(db *sql.DB, p Netflix_shows) error {
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
		log.Printf("%s when inserting row into netflix_show_details table", err)
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

func InsertShows(db *sql.DB, netflix_titles [][]string) error {
	var err error
	for _, line := range netflix_titles {
		rec := Netflix_shows{
			Show_id:      line[0],
			Show_type:    line[1],
			Title:        line[2],
			Director:     line[3],
			Cast:         line[4],
			Country:      line[5],
			Date_added:   line[6],
			Release_year: line[7],
			Rating:       line[8],
			Duration:     line[9],
			Listed_in:    line[10],
			Description:  line[11],
		}
		err = InsertRows(db, rec)
		if err != nil {
			log.Printf("Insert failed with error %s", err)
		}
	}
	log.Printf("Successfully inserted rows in table")
	return err
}

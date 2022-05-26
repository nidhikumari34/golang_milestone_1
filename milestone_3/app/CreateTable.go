package app

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//creating netflix_show_details table
func CreateTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS netflix_show_details_api(show_id VARCHAR(500) PRIMARY KEY,
															show_type VARCHAR(500),
															title VARCHAR(500),
															director VARCHAR(500),
															cast VARCHAR(5000),
															country VARCHAR(500),
															date_added VARCHAR(500),
															release_year VARCHAR(500),
															rating VARCHAR(500),
															duration VARCHAR(500),
															listed_in VARCHAR(500),
															description VARCHAR(5000))`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("%s when creating netflix_show_details table", err)
		return err
	}
	log.Printf("Successfully created table")
	return nil
}

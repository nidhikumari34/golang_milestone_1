package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type netflix_shows_db struct {
	show_id      string
	show_type    string
	title        string
	director     string
	cast         string
	country      string
	date_added   string
	release_year string
	rating       string
	duration     string
	listed_in    string
	description  string
}

const (
	username = "root"
	password = "knightrider@6N"
	hostname = "127.0.0.1:3306"
	dbname   = "golang"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

//connecting to MySQL DB
func dbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("%s when opening DB\n", err)
		return nil, err
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("%s when creating DB\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("%s when fetching rows", err)
		return nil, err
	}
	log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("%s when opening DB", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	return db, nil
}

//creating netflix_show_details table
func create_Table(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS netflix_show_details(show_id VARCHAR(500) PRIMARY KEY,
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
	return nil
}

//inserting rows in netflix_show_details table
func insert(db *sql.DB, p netflix_shows_db) error {
	query := "INSERT INTO netflix_show_details(show_id,show_type,title,director,cast,country,date_added,release_year,rating,duration,listed_in,description) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("%s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx,
		p.show_id,
		p.show_type,
		p.title,
		p.director,
		p.cast, p.country,
		p.date_added,
		p.release_year,
		p.rating,
		p.duration,
		p.listed_in,
		p.description)
	if err != nil {
		log.Printf("%s when inserting row into netflix_show_details table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("%s when finding rows affected", err)
		return err
	}
	log.Printf("%d netflix_show_details created: show_id = %s ", rows, p.show_id)
	return nil
}

func insert_db(netflix_titles [][]string) {
	db, err := dbConnection()
	if err != nil {
		log.Printf("%s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")

	err = create_Table(db)
	if err != nil {
		log.Printf("Create netflix_show_details table failed with error %s", err)
		return
	}

	for _, line := range netflix_titles {
		rec := netflix_shows_db{
			show_id:      line[0],
			show_type:    line[1],
			title:        line[2],
			director:     line[3],
			cast:         line[4],
			country:      line[5],
			date_added:   line[6],
			release_year: line[7],
			rating:       line[8],
			duration:     line[9],
			listed_in:    line[10],
			description:  line[11],
		}
		err = insert(db, rec)
		if err != nil {
			log.Printf("Insert failed with error %s", err)
			return
		}
	}
}

package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type netflixshows struct {
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

func main() {
	csvFile, err := os.Open("netflix_titles.csv")
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	netflix_titles, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Println(err)
	}

	netflix_titles = sort_csv(netflix_titles)

	db, err := sql.Open("mysql", "root:<knightrider@6N>@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO netflix_show_details VALUES('1')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Success!")

	//insert_db(db, netflixshows)
	//tv_shows(netflix_titles)
	//horror_movies(netflix_titles)
	//indian_movies(netflix_titles)
}

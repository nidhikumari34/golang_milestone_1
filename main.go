package main

import (
	"encoding/csv"
	"log"
	"os"
)

type netflix_shows struct {
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

	insert_db(netflix_titles)
	netflix_titles = sort_csv(netflix_titles)
	tv_shows(netflix_titles)
	horror_movies(netflix_titles)
	indian_movies(netflix_titles)
}

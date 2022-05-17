package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

//First n records where type is TV Show
func tv_shows(netflix_titles [][]string) {
	var n int
	fmt.Printf("\nEnter the value of n for TV Shows: ")
	fmt.Scanf("%d", &n)
	i := 1
	for _, line := range netflix_titles {
		rec := netflix_shows{
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
		if rec.show_type == "TV Show" && i <= n {
			fmt.Println(rec.show_id + ", " + rec.show_type + ", " + rec.title + ", " + rec.director + ", " + rec.cast + ", " + rec.country + ", " + rec.date_added + ", " + rec.release_year + ", " + rec.rating + ", " + rec.duration + ", " + rec.listed_in + ", " + rec.description)
			fmt.Println()
			i++
		}
		if i > n {
			break
		}
	}
}

//First n records where listed_in have Horror Movies
func horror_movies(netflix_titles [][]string) {
	var n int
	fmt.Printf("\nEnter the value of n for Horror Movies: ")
	fmt.Scanf("%d", &n)
	i := 1
	for _, line := range netflix_titles {
		rec := netflix_shows{
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
		if strings.Contains(rec.listed_in, "Horror Movies") && i <= n {
			fmt.Println(rec.show_id + ", " + rec.show_type + ", " + rec.title + ", " + rec.director + ", " + rec.cast + ", " + rec.country + ", " + rec.date_added + ", " + rec.release_year + ", " + rec.rating + ", " + rec.duration + ", " + rec.listed_in + ", " + rec.description)
			fmt.Println()
			i++
		}
		if i > n {
			break
		}
	}
}

//First n records where type is Movie and country is India
func indian_movies(netflix_titles [][]string) {
	var n int
	fmt.Printf("\nEnter the value of n for Indian Movies: ")
	fmt.Scanf("%d", &n)
	i := 1
	for _, line := range netflix_titles {
		rec := netflix_shows{
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
		if rec.show_type == "Movie" && rec.country == "India" && i <= n {
			fmt.Println(rec.show_id + ", " + rec.show_type + ", " + rec.title + ", " + rec.director + ", " + rec.cast + ", " + rec.country + ", " + rec.date_added + ", " + rec.release_year + ", " + rec.rating + ", " + rec.duration + ", " + rec.listed_in + ", " + rec.description)
			fmt.Println()
			i++
		}
		if i > n {
			break
		}
	}
}

func main() {
	csvFile, err := os.Open("netflix_titles.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	netflix_titles, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	/*for _, line := range netflix_titles {
		rec := netflix_shows{
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
		fmt.Println("\n" + rec.listed_in)
	}*/

	tv_shows(netflix_titles)
	horror_movies(netflix_titles)
	indian_movies(netflix_titles)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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

//First n records b/w start and end date where type is TV Show
func tv_shows(netflix_titles [][]string) {
	var n int
	var start, end string
	var start_dt, end_dt time.Time
	i := 1
	reader := bufio.NewReader(os.Stdin)

	log.Printf("Enter the value of n for TV Shows :")
	fmt.Scanf("%d", &n)

	log.Println("Enter start date (Month DD, YYYY) :")
	start, _ = reader.ReadString('\n')

	log.Println("Enter end date (Month DD, YYYY) :")
	end, _ = reader.ReadString('\n')

	start_dt, _ = time.Parse("January 2, 2006", strings.TrimSpace(start))
	end_dt, _ = time.Parse("January 2, 2006", strings.TrimSpace(end))
	startTime := time.Now()

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
		date_added, _ := time.Parse("January 02, 2006", rec.date_added)

		if rec.show_type == "TV Show" && i <= n && date_added.Before(end_dt) && date_added.After(start_dt) {
			fmt.Println("\n", rec.show_id+", "+rec.show_type+", "+rec.title+", "+rec.director+", "+rec.cast+", "+rec.country+", "+rec.date_added+", "+rec.release_year+", "+rec.rating+", "+rec.duration+", "+rec.listed_in+", "+rec.description)
			i++
		}
		if i > n {
			break
		}
	}
	fmt.Println()
	log.Println("Execution time taken:", time.Now().Sub(startTime))
}

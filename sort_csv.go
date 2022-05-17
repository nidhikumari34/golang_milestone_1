package main

import "sort"

type netflix_shows3 struct {
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

//sorting by duration
func sort_csv(netflix_titles [][]string) [][]string {
	sort.Slice(netflix_titles, func(i, j int) bool {
		return netflix_titles[i][9] < netflix_titles[j][9]
	})
	return netflix_titles
}

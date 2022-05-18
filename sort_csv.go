package main

import "sort"

//sorting by duration
func sort_csv(netflix_titles [][]string) [][]string {
	sort.Slice(netflix_titles, func(i, j int) bool {
		return netflix_titles[i][9] < netflix_titles[j][9]
	})
	return netflix_titles
}

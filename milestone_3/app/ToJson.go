package app

import (
	_ "github.com/go-sql-driver/mysql"
)

func ToJson() []Netflix_shows {
	var tvshows []Netflix_shows

	netflix_titles := ReadCSV()
	for _, line := range netflix_titles[1:] {
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
		tvshows = append(tvshows, rec)
	}
	return tvshows
}

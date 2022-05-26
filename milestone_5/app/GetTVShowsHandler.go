package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

//First n TV Shows
func GetTVShows(w http.ResponseWriter, r *http.Request) {
	var tvshows []Netflix_shows
	var count_val = r.URL.Query().Get("count")
	var n, err = strconv.Atoi(count_val)
	if err != nil {
		log.Printf("Invalid input")
	}

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
	EndTime = time.Now()
	ExecTime = EndTime.Sub(StartTime)
	w.Header().Add("X-TIME-TO-EXECUTE", ExecTime.String())
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tvshows[:n])
}

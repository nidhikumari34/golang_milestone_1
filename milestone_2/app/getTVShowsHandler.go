package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//First n TV Shows
func getTVShows(w http.ResponseWriter, r *http.Request) {
	var tvshows []Netflix_shows
	vars := mux.Vars(r)
	count_val, err := strconv.Atoi(vars["count"])
	if err != nil {
		log.Printf("Invalid input")
	}

	netflix_titles := Read_csv()
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
	endTime = time.Now()
	execTime = endTime.Sub(startTime)
	w.Header().Add("TIME-TO-EXECUTE", execTime.String())
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tvshows[:count_val])
}

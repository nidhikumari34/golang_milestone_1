package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//Shows between start and end date
func GetBetweenDates(w http.ResponseWriter, r *http.Request) {
	var betweenShows []Netflix_shows
	var start_date = r.URL.Query().Get("startDate")
	var end_date = r.URL.Query().Get("endDate")

	start_dt, _ := time.Parse("January 2, 2006", start_date)
	end_dt, _ := time.Parse("January 2, 2006", end_date)
	log.Println(start_dt)
	log.Println(end_dt)
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
		date_added, _ := time.Parse("January 02, 2006", rec.Date_added)
		if date_added.Before(end_dt) && date_added.After(start_dt) {
			betweenShows = append(betweenShows, rec)
		}
	}
	if betweenShows != nil {
		EndTime = time.Now()
		ExecTime = EndTime.Sub(StartTime)
		w.Header().Add("X-TIME-TO-EXECUTE", ExecTime.String())
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(betweenShows)
	} else {
		log.Printf("No shows or Invalid input")
	}
}

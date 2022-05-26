package app

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//Indian Movies
func GetCountryMovies(w http.ResponseWriter, r *http.Request) {
	var countryMovies []Netflix_shows
	var country_val = r.URL.Query().Get("country")

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
		if rec.Country == country_val {
			countryMovies = append(countryMovies, rec)
		}
	}
	if countryMovies != nil {
		EndTime = time.Now()
		ExecTime = EndTime.Sub(StartTime)
		w.Header().Add("X-TIME-TO-EXECUTE", ExecTime.String())
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(countryMovies)
	} else {
		log.Printf("Invalid input")
	}
}

package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

//Horror Movies
func GetMovieType(w http.ResponseWriter, r *http.Request) {
	var movie_type []Netflix_shows
	vars := mux.Vars(r)
	type_val := vars["movieType"]

	if type_val == "horror" {
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
			if strings.Contains(rec.Listed_in, "Horror Movies") {
				movie_type = append(movie_type, rec)
			}
		}
		EndTime = time.Now()
		ExecTime = EndTime.Sub(StartTime)
		w.Header().Add("X-TIME-TO-EXECUTE", ExecTime.String())
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movie_type)
	} else {
		log.Printf("Invalid input")
	}
}

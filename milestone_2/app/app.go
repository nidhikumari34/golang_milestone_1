package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var startTime, endTime time.Time
var execTime time.Duration

//struct for show details
type Netflix_shows struct {
	Show_id      string `json:"show_id"`
	Show_type    string `json:"show_type"`
	Title        string `json:"title"`
	Director     string `json:"director"`
	Cast         string `json:"cast"`
	Country      string `json:"country"`
	Date_added   string `json:"date_added"`
	Release_year string `json:"release_year"`
	Rating       string `json:"rating"`
	Duration     string `json:"duration"`
	Listed_in    string `json:"listed_in"`
	Description  string `json:"description"`
}

func Start() {
	router := mux.NewRouter()

	//define route
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(getTVShows))).Methods(http.MethodGet).Queries("count", "{count:[0-9]+}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(getMovieType))).Methods(http.MethodGet).Queries("movieType", "{movieType}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(getCountryMovies))).Methods(http.MethodGet).Queries("country", "{country}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(getBetweenDates))).Methods(http.MethodGet).Queries("startDate", "{startDate}", "endDate", "{endDate}")

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

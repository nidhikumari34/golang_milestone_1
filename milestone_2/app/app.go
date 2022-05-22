package app

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var Key = []byte("abc")
var TokenString string
var StartTime, EndTime time.Time
var ExecTime time.Duration

var Users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(GetTVShows))).Methods(http.MethodGet).Queries("count", "{count:[0-9]+}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(GetMovieType))).Methods(http.MethodGet).Queries("movieType", "{movieType}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(GetCountryMovies))).Methods(http.MethodGet).Queries("country", "{country}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(GetBetweenDates))).Methods(http.MethodGet).Queries("startDate", "{startDate}", "endDate", "{endDate}")

	router.HandleFunc("/login", Login).Methods("POST")

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

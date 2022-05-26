package app

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"gopkg.in/robfig/cron.v2"
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
	//cron job to schedule syncing of the database with every new netflix show details in the file
	c := cron.New()
	c.AddFunc("@every 1m", DBSyncCronJob)
	c.Start()

	router := mux.NewRouter()

	//define route
	router.HandleFunc("/insertDB", InsertCSVIntoDB).Methods("POST")      //insert CSV data into DB
	router.HandleFunc("/insertDBJSON", InsertJSONIntoDB).Methods("POST") //insert JSON data from POST request to DB
	router.HandleFunc("/getData", GetData).Methods(http.MethodGet).Queries("DataSource", "{DataSource}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(GetTVShows))).Methods(http.MethodGet).Queries("count", "{count:[0-9]+}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(GetMovieType))).Methods(http.MethodGet).Queries("movieType", "{movieType}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(GetCountryMovies))).Methods(http.MethodGet).Queries("country", "{country}")
	router.HandleFunc("/tvshows", AuthTokenMiddleware(TimingMiddleware(GetBetweenDates))).Methods(http.MethodGet).Queries("startDate", "{startDate}", "endDate", "{endDate}")

	router.HandleFunc("/login", Login).Methods("POST")

	//starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}

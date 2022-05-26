package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/robfig/cron.v2"
)

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

	//starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}

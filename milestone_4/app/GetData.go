package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//fetch shows from data source as csv or db
func GetData(w http.ResponseWriter, r *http.Request) {
	var data_source = r.URL.Query().Get("DataSource")

	if strings.ToUpper(data_source) == "DB" {
		Netflix_shows := ReadDataFromDB()
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Netflix_shows)
	} else if strings.ToUpper(data_source) == "CSV" {
		Netflix_shows := ReadDataFromCSV()
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Netflix_shows)
	} else {
		log.Println("Invalid Data Source")
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Data Source"))
		return
	}

}

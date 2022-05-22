package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//fetch shows from data source as csv or db
func GetData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data_souce := vars["DataSource"]

	if strings.ToUpper(data_souce) == "DB" {
		Netflix_shows := ReadDataFromDB()
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Netflix_shows)
	} else if strings.ToUpper(data_souce) == "CSV" {
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

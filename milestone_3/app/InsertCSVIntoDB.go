package app

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var netflix_titles [][]string

const (
	username = "root"
	password = "knightrider@6N"
	hostname = "127.0.0.1:3306"
	dbname   = "golang_api"
)

func InsertCSVIntoDB(w http.ResponseWriter, r *http.Request) {
	var err error
	netflix_titles = ReadCSV()

	db, err = DBConnection()
	if err != nil {
		log.Printf("%s when getting db connection", err)
		return
	}
	//defer db.Close()
	log.Printf("Successfully connected to database")

	err = CreateTable(db)
	if err != nil {
		log.Printf("Create netflix_show_details table failed with error %s", err)
		return
	}

	//DBSync()

	err = InsertShows(db, netflix_titles)
	if err != nil {
		log.Printf("Inserting netflix_show_details failed for few or all records")
		return
	}

	//fmt.Println(json_data)
	/*payloadBuf := new(bytes.Buffer)
	//json.NewEncoder(payloadBuf).Encode(json_data)

	err := json.NewDecoder(payloadBuf).Decode(&json_data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/
	//fmt.Println(json_data)
	/*resp, err := http.NewRequest("POST", "http://localhost:8000/post", payloadBuf)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()*/
	//w.Header().Add("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(json_data)
	//json.NewEncoder(r.Body).Encode(json_data)
	//json.NewDecoder(r.Body).Decode(&json_data)
	/*
		db, er := DBConnection()
		if er != nil {
			log.Printf("%s while connecting to DB", er)
			//return err
		}
		er = CreateTable(db)
		if er != nil {
			log.Printf("%s when creating netflix_show_details table", er)
			//return err
		}
		//fmt.Println("%T", json_data)
		//InsertRows(db, json_data)*/

}

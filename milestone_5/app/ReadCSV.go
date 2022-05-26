package app

import (
	"encoding/csv"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//Read csv and return records
func ReadCSV() [][]string {
	csvFile, err := os.Open("netflix_titles.csv")
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	netflix_titles, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		log.Println("Error while reading CSV: ", err)
	}
	return netflix_titles
}

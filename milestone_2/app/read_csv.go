package app

import (
	"encoding/csv"
	"log"
	"os"
)

//Read csv and return records as strings
func Read_csv() [][]string {
	csvFile, err := os.Open("netflix_titles.csv")
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	netflix_titles, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		log.Println(err)
	}
	return netflix_titles
}

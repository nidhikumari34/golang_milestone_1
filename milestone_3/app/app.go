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
	c.AddFunc("@every 5s", DBSyncCronJob)
	c.Start()

	router := mux.NewRouter()

	//define route
	router.HandleFunc("/insertDB", InsertCSVIntoDB).Methods("POST")
	router.HandleFunc("/getData", GetData).Methods(http.MethodGet).Queries("DataSource", "{DataSource}")

	//starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

/*func ReadCSVFromHttpRequest(req *http.Request) ([][]string, error) {
	// parse POST body as csv
	reader := csv.NewReader(req.Body)
	var results [][]string
	for {
		// read one row from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// add record to result set
		results = append(results, record)
	}
	return results, nil
}
*/

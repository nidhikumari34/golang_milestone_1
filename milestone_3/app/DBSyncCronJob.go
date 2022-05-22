package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//automatic cron job for DB sync
func DBSyncCronJob() {
	var err error
	db, err = DBConnection()
	if err != nil {
		log.Printf("%s when getting db connection", err)
		return
	}
	log.Printf("Successfully connected to database")

	err = CreateTable(db)
	if err != nil {
		log.Printf("Create netflix_show_details table failed with error %s", err)
		return
	}
	netflix_titles = ReadCSV()
	InsertShows(db, netflix_titles)
}

//cron job for DB sync after API POST call
/*func DBSync() {
	gocron.Start()
	//ch := gocron.Start()
	//go test(ch)

	gocron.Every(5).Seconds().Do(task)

	time.Sleep(40 * time.Second)

	gocron.Clear()
	fmt.Println("All task removed")
}

//func test(stop chan bool) {
//	time.Sleep(20 * time.Second)
//	gocron.Clear()
//	fmt.Println("All task removed")
//	close(stop)
//}

func task() {
	netflix_titles = ReadCSV()
	InsertShows(db, netflix_titles)
}*/

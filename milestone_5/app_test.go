package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"show/app"
	"testing"
)

var auth_token string

//correct login details
func TestLoginPass(t *testing.T) {
	var jsonStr = []byte(`{"username": "user1","password": "password1"}`)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.Login)
	handler.ServeHTTP(rr, req)

	auth_token = rr.Body.String()

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := 0
	if len(rr.Body.String()) == expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//incorrect login details
func TestLoginFail(t *testing.T) {
	var jsonStr = []byte(`{"username": "user1","password": "password"}`)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.Login)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := 0
	if len(rr.Body.String()) == expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

//missing login details
func TestLoginMissing(t *testing.T) {
	var jsonStr = []byte(``)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.Login)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

//missing token for n shows
func TestGetTVShowsTokenMissing(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("count", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetTVShows)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}

//invalid token for n shows
func TestGetTVShowsTokenInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("count", "1")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-Auth-Token", "test")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetTVShows)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}
}

//correct token for n shows
func TestGetTVShows(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("count", "1")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-Auth-Token", app.TokenString)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetTVShows)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//correct token for n shows
func TestGetTVShowsInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("count", "test")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-Auth-Token", app.TokenString)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetTVShows)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//correct token for Horror Movies
func TestGetMovieTypeHorror(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("movieType", "horror")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-Auth-Token", app.TokenString)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetMovieType)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//correct token for Horror Movies
func TestGetMovieTypeNotHorror(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("movieType", "")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-Auth-Token", app.TokenString)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetMovieType)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//correct token for Indian Movies
func TestGetCountryMovies(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("country", "India")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-Auth-Token", app.TokenString)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetCountryMovies)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//correct token for Indian Movies
func TestGetCountryNoMovies(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("country", "")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-Auth-Token", app.TokenString)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetCountryMovies)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//correct token for shows between start and end date
func TestGetBetweenDates(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("startDate", "August 14, 2020")
	q.Add("endDate", "August 14, 2022")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-Auth-Token", app.TokenString)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetBetweenDates)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//correct token for shows between start and end date
func TestGetBetweenDatesNoShow(t *testing.T) {
	req, err := http.NewRequest("GET", "/tvshows", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("startDate", "test")
	q.Add("endDate", "test")
	req.URL.RawQuery = q.Encode()
	req.Header.Add("X-Auth-Token", app.TokenString)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AuthTokenMiddleware(app.TimingMiddleware(app.GetBetweenDates)))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//get data from source CSV
func TestGetData(t *testing.T) {
	req, err := http.NewRequest("GET", "/getData", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("DataSource", "csv")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetData)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK && rr.Body != nil {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//get data from source DB
func TestGetDataDB(t *testing.T) {
	req, err := http.NewRequest("GET", "/getData", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("DataSource", "db")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetData)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK && rr.Body != nil {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//get data from invalid source
func TestGetDataInvalid(t *testing.T) {
	req, err := http.NewRequest("GET", "/getData", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("DataSource", "test")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetData)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest && rr.Body != nil {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

//insert CSV data into DB
func TestInsertCSVIntoDB(t *testing.T) {
	req, err := http.NewRequest("POST", "/insertDB", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.InsertCSVIntoDB)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//insert JSON data from POST request to DB
func TestInsertJSONIntoDB(t *testing.T) {
	var Str = []byte(`[{"show_id":"uyf","show_type":"TV Show","title":"3%","director":"","cast":"João Miguel, Bianca Comparato, Michel Gomes, Rodolfo Valente, Vaneza Oliveira, Rafael Lozano, Viviane Porto, Mel Fronckowiak, Sergio Mamberti, Zezé Motta, Celso Frateschi","country":"Brazil","date_added":"August 14, 2020","release_year":"2020","rating":"TV-MA","duration":"4 Seasons","listed_in":"International TV Shows, TV Dramas, TV Sci-Fi & Fantasy","description":"In a future where the elite inhabit an island paradise far from the crowded slums, you get one chance to join the 3% saved from squalor."}]`)
	req, err := http.NewRequest("POST", "/insertDB", bytes.NewBuffer(Str))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.InsertJSONIntoDB)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

//table creation
func TestCreateTable(t *testing.T) {
	db, _ := app.DBConnection()
	err := app.CreateTable(db)
	if err != nil {
		t.Errorf("failed to create table")
	}
}

//db connection
func TestDBConnection(t *testing.T) {
	_, err := app.DBConnection()

	if err != nil {
		t.Errorf("failed to create db connection")
	}
}

//cron job
func TestDBSyncCronJob(t *testing.T) {
	app.DBSyncCronJob()
}

//read csv
func TestReadCSV(t *testing.T) {
	var str = app.ReadCSV()
	if len(str) == 0 {
		t.Errorf("failed to read csv")
	}
}

//invalid passed token
func TestInvalidValidToken(t *testing.T) {
	var match = app.ValidToken("test")
	if match {
		t.Errorf("failed to read csv")
	}
}

package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

// A jsonErr object includes the return code and error message and
// is returned when an error condition occurs.
type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// A carquery provides the manufacturer, minimum MPG, and
// maximum price to use when searching for a car.
type carquery struct {
	Manufacturer string
	Minmpg       string
	Maxprice     string
}

// parseSqliteResultsForCars queries the SQLite3 database using
// the query string and parameters passed as arguments and returns
// an array of Car that meet the search criteria.
func parseSqliteResultsForCars(qs string, params ...string) Cars {
	var carid int
	var price int
	var mpg int
	var manufacturer string
	var model string
	var enginesize int
	var horsepower int
	var wheelbase int
	var passengers int
	var cars Cars
	var rows *sql.Rows

	db, err := sql.Open("sqlite3", "./cars93.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch len(params) {
	case 0:
		rows, err = db.Query(qs)
	case 1:
		rows, err = db.Query(qs, params[0])
	case 2:
		rows, err = db.Query(qs, params[0], params[1])
	case 3:
		rows, err = db.Query(qs, params[0], params[1], params[2])
	case 4:
		rows, err = db.Query(qs, params[0], params[1], params[2], params[3])
	case 5:
		rows, err = db.Query(qs, params[0], params[1], params[2], params[3], params[4])
	case 6:
		rows, err = db.Query(qs, params[0], params[1], params[2], params[3], params[4], params[5])
	case 7:
		rows, err = db.Query(qs, params[0], params[1], params[2], params[3], params[4], params[5], params[6])
	case 8:
		rows, err = db.Query(qs, params[0], params[1], params[2], params[3], params[4], params[5], params[6], params[7])
	case 9:
		rows, err = db.Query(qs, params[0], params[1], params[2], params[3], params[4], params[5], params[6], params[7], params[8])

	default:
		log.Fatal("Number of params exceeded 9")
	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&carid, &manufacturer, &model, &price, &mpg, &enginesize, &horsepower, &wheelbase, &passengers)
		if err != nil {
			log.Fatal(err)
		}

		cars = append(cars, Car{
			ID:           carid,
			Manufacturer: manufacturer,
			Model:        model,
			Price:        price,
			MPG:          mpg,
			Enginesize:   enginesize,
			Horsepower:   horsepower,
			Wheelbase:    wheelbase,
			Passengers:   passengers})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return cars
}

// parseSqliteResultsForArray queries the SQLite3 database using
// the query string and parameters passed as arguments and returns
// an array of string that meet the search criteria.
func parseSqliteResultsForArray(qs string, params ...string) []string {
	var results []string
	var token string
	var rows *sql.Rows

	db, err := sql.Open("sqlite3", "./cars93.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	switch len(params) {
	case 0:
		rows, err = db.Query(qs)
	case 1:
		rows, err = db.Query(qs, params[0])
	case 2:
		rows, err = db.Query(qs, params[0], params[1])
	case 3:
		rows, err = db.Query(qs, params[0], params[1], params[2])
	case 4:
		rows, err = db.Query(qs, params[0], params[1], params[2], params[3])
	case 5:
		rows, err = db.Query(qs, params[0], params[1], params[2], params[3], params[4])
	default:
		log.Fatal("Number of params exceeded 5")
	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&token)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, token)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return results
}

// slashRedirect redirects to the home page
func slashRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://danwritesandcodes.com", 301)
}

// findAllCars creates a SQL SELECT statement to get
// all cars in the database and returns an array of JSON-encoded
// Cars.
func findAllCars(w http.ResponseWriter, r *http.Request) {
	var cars Cars
	query := `select id,manufacturer,model,price,mpg_highway,enginesize,horsepower,wheelbase,passengers from cars;`
	cars = parseSqliteResultsForCars(query)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		panic(err)
	}
}

// findAllCarManufacturers creates a SQL SELECT statement to get
// all car manufacturers in the database and returns an array of
// strings.
func findAllCarManufacturers(w http.ResponseWriter, r *http.Request) {
	query := `select distinct(manufacturer) from cars;`
	manufacturers := parseSqliteResultsForArray(query)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(manufacturers); err != nil {
		panic(err)
	}
}

// getCarsByMaxPrice creates a SQL SELECT statement to get
// all cars less than or equal to the specified maximum price
// and returns an array of JSON-encoded Cars.
func getCarsByMaxPrice(w http.ResponseWriter, r *http.Request) {
	var cars Cars
	vars := mux.Vars(r)

	query := `select id,manufacturer,model,price,mpg_highway,enginesize,horsepower,wheelbase,passengers from cars where price <= ?`
	cars = parseSqliteResultsForCars(query, vars["maxPrice"])

	if len(cars) == 0 {
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		panic(err)
	}
}

// getCarsByMinMPG creates a SQL SELECT statement to get
// all cars less than or equal to the specified highway miles per gallon (MPG)
// and returns an array of JSON-encoded Cars.
func getCarsByMinMPG(w http.ResponseWriter, r *http.Request) {
	var cars Cars
	vars := mux.Vars(r)

	query := `select id,manufacturer,model,price,mpg_highway,enginesize,horsepower,wheelbase,passengers from cars where mpg_highway >= ?`

	// results := getSqliteResults(query)
	cars = parseSqliteResultsForCars(query, vars["minMPG"])

	if len(cars) == 0 {
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		panic(err)
	}
}

// getCarsByID creates a SQL SELECT statement to get
// the car that matches the specified id
// and returns an array of a single JSON-encoded Car.
func getCarsByID(w http.ResponseWriter, r *http.Request) {
	var cars Cars
	vars := mux.Vars(r)

	query := `select id,manufacturer,model,price,mpg_highway,enginesize,horsepower,wheelbase,passengers from cars where id = ?`

	cars = parseSqliteResultsForCars(query, vars["id"])

	if len(cars) == 0 {
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		panic(err)
	}
}

// getCarsByManufacturer creates a SQL SELECT statement to get
// all cars of the specified manufacturer
// and returns an array of JSON-encoded Cars.
func getCarsByManufacturer(w http.ResponseWriter, r *http.Request) {
	var cars Cars
	vars := mux.Vars(r)

	query := `select id,manufacturer,model,price,mpg_highway,enginesize,horsepower,wheelbase,passengers from cars where manufacturer = ?`

	cars = parseSqliteResultsForCars(query, vars["carManufacturer"])

	if len(cars) == 0 {
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		panic(err)
	}
}

// getCarsByCriteria creates a SQL SELECT statement to get
// all cars that match the minimum highway MPG, maximum price, and
// specified manufacturer passed as a JSON object.
// An array of JSON-encoded Cars is returned.
func getCarsByCriteria(w http.ResponseWriter, r *http.Request) {
	var cars Cars
	var newcarquery carquery
	var query string
	maxPrice := "100000000000"
	minMPG := "0"

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newcarquery)
	if err != nil {
		log.Printf("Error decoding newcarquery (%v)", err)
		return
	}
	if len(newcarquery.Minmpg) > 0 {
		minMPG = newcarquery.Minmpg
	}
	if len(newcarquery.Maxprice) > 0 {
		maxPrice = newcarquery.Maxprice
	}
	if len(newcarquery.Manufacturer) > 0 {
		query = `select id,manufacturer,model,price,mpg_highway,enginesize,horsepower,wheelbase,passengers from cars where manufacturer = ? and mpg_highway >= ? and price <= ?`
		cars = parseSqliteResultsForCars(query, newcarquery.Manufacturer, minMPG, maxPrice)
	} else {
		query = `select id,manufacturer,model,price,mpg_highway,enginesize,horsepower,wheelbase,passengers from cars where mpg_highway >= ? and price <= ?`
		cars = parseSqliteResultsForCars(query, minMPG, maxPrice)
	}

	if len(cars) == 0 {
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		panic(err)
	}
}

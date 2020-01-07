/*
Package main implements a simple server for the cars93 data set.

Supported methods are:

http://localhost:8080/cars/allcars				 : Return all cars in the database

http://localhost:8080/cars/id/{id}				 : Return car with id = {id}

http://localhost:8080/cars/allmanufacturers		 : Return all car manufacturers

http://localhost:8080/cars/manufacturer/{mfgr}	 : Return all cars manufactured by {mfgr}

http://localhost:8080/cars/maxprice/{price}		 : Return all cars with a price less than or equal to {price}

http://localhost:8080/cars/minmpg/{mpg}			 : Return all cars with highway MPG greater than or equal to {mpg}

http://localhost:8080/cars/criteria				 : Return all cars that match the manufacturer, MPG and price criteria specified as a JSON object

*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running on port :8080 ...")
	fmt.Println("  Try http://localhost:8080/cars")
	fmt.Println("      http://localhost:8080/cars/1")

	router := newRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

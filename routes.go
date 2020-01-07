package main

import "net/http"

// A Route represents the information needed to fulfill a route request
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"slashRedirect",
		"GET",
		"/",
		slashRedirect,
	},
	Route{
		"findAllCars",
		"GET",
		"/cars",
		findAllCars,
	},
	Route{
		"findAllCars",
		"GET",
		"/cars/allcars",
		findAllCars,
	},
	Route{
		"findAllCarManufacturers",
		"GET",
		"/cars/allmanufacturers",
		findAllCarManufacturers,
	},
	Route{
		"getCarsByManufacturer",
		"GET",
		"/cars/manufacturer/{carManufacturer}",
		getCarsByManufacturer,
	},
	Route{
		"getCarsByMaxPrice",
		"GET",
		"/cars/maxprice/{maxPrice}",
		getCarsByMaxPrice,
	},
	Route{
		"getCarsByMinMPG",
		"GET",
		"/cars/minmpg/{minMPG}",
		getCarsByMinMPG,
	},
	Route{
		"getCarsByID",
		"GET",
		"/cars/id/{id}",
		getCarsByID,
	},
	Route{
		"getCarsByCriteria",
		"POST",
		"/cars/criteria",
		getCarsByCriteria,
	},
}

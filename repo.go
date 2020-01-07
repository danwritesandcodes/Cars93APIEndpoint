package main

// A Car contains an ID, a manufacturer, a model, a price, and a (highway) MPG
type Car struct {
	ID           int    `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Price        int    `json:"price"`
	MPG          int    `json:"mpg"`
	Enginesize   int    `json:"enginesize"`
	Horsepower   int    `json:"horsepower"`
	Wheelbase    int    `json:"wheelbase"`
	Passengers   int    `json:"passengers"`
}

type Cars []Car

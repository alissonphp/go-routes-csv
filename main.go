package main

import (
	"go-best-route/adapters/csv"
	"go-best-route/application"
)

func main() {
	c := csv.RoutesCSV{}
	r := application.Route{From: "SLZ", To: "BSB", Price: 90}
	c.Save(&r)
}

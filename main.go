package main

import (
	"fmt"
	"go-best-route/adapters/csv"
)

func main() {
	c := csv.RoutesCSV{}
	fmt.Println(c.ReadAll())
}

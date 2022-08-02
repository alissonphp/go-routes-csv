package csv

import (
	"encoding/csv"
	"go-best-route/application"
	"log"
	"os"
	"strconv"
)

type RoutesCSV struct {
}

func (r *RoutesCSV) ReadAll() ([]application.Route, error) {
	f, err := os.Open("input-routes.csv")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return []application.Route{}, err
	}

	var routes []application.Route

	for _, line := range reader {
		s, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		route := application.Route{
			From:  line[0],
			To:    line[1],
			Price: s,
		}

		routes = append(routes, route)
	}

	return routes, err
}

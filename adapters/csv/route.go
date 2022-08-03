package csv

import (
	"encoding/csv"
	"go-best-route/application"
	"log"
	"os"
	"strconv"
)

type RoutesCSV struct {
	csvFile string
}

func NewReadCsvFile(path string) *RoutesCSV {
	return &RoutesCSV{csvFile: path}
}

func (r *RoutesCSV) Get(from string, to string) (application.RouteInterface, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RoutesCSV) List() ([]application.RouteInterface, error) {
	f, err := os.Open(r.csvFile)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	var routes []application.RouteInterface

	for _, line := range reader {
		s, err := strconv.Atoi(line[2])
		if err != nil {
			log.Fatal(err)
		}
		route := application.Route{
			From:  line[0],
			To:    line[1],
			Price: s,
		}

		routes = append(routes, &route)
	}

	return routes, err
}

func (r *RoutesCSV) Save(route application.RouteInterface) (application.RouteInterface, error) {
	f, err := os.OpenFile(r.csvFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(f)
	w.Write([]string{
		route.GetFrom(), route.GetTo(), strconv.Itoa(route.GetPrice()),
	})
	w.Flush()

	return route, nil
}

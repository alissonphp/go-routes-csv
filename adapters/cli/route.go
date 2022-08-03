package cli

import (
	"fmt"
	"go-best-route/application"
)

func Run(service application.RouteServiceInterface, from string, to string) (string, error) {
	var result = ""
	res, err := service.SearchBest(from, to)

	if err != nil {
		return result, err
	}

	result = fmt.Sprintf("best route: %s > $%d", res.GetFlyPath(), res.GetTotalCost())

	return result, nil
}

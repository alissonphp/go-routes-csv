package cli

import "go-best-route/application"

func Run(service application.RouteServiceInterface, path string, from string, to string) (application.BestRoute, error) {
	var result application.BestRoute
	res, err := service.SearchBest(from, to)

	if err != nil {
		return result, err
	}

	return res, nil
}

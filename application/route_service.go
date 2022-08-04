package application

import (
	"fmt"
	"sort"
)

type RouteService struct {
	Persistence RoutePersistenceInterface
}

func NewRouteService(persistence RoutePersistenceInterface) *RouteService {
	return &RouteService{Persistence: persistence}
}

func (s *RouteService) Get(from string, to string) (RouteInterface, error) {
	route, err := s.Persistence.Get(from, to)
	if err != nil {
		return nil, err
	}
	return route, nil

}

func (s *RouteService) List() ([]Route, error) {
	routes, err := s.Persistence.List()
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func (s *RouteService) Save(from string, to string, price int) (RouteInterface, error) {
	route := NewRoute()
	route.From = from
	route.To = to
	route.Price = price

	res, err := s.Persistence.Save(route)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *RouteService) SearchBest(from string, to string) (BestRoute, error) {

	var possibilities []Route

	routes, _ := s.List()
	for _, route := range routes {
		if route.GetFrom() == from {
			possibilities = append(possibilities, route)
		}
	}

	bs := s.checkPath(possibilities, to)
	sort.Sort(ByCost(bs))
	return bs[0], nil
}

func (s *RouteService) checkPath(routes []Route, final string) []BestRoute {
	var bests []BestRoute
	for _, route := range routes {
		b := BestRoute{
			FlyPath:   fmt.Sprintf("%s-%s", route.GetFrom(), route.GetTo()),
			TotalCost: route.GetPrice(),
		}

		s.getNextPoint(&b, route.GetTo(), final)

		bests = append(bests, b)
	}

	return bests
}

func (s *RouteService) getNextPoint(path *BestRoute, next string, final string) {
	routes, _ := s.List()
	for _, route := range routes {
		if route.GetFrom() == next {
			path.FlyPath = fmt.Sprintf("%s-%s", path.GetFlyPath(), route.GetTo())
			path.TotalCost = path.GetTotalCost() + route.GetPrice()
			if route.GetTo() != final {
				s.getNextPoint(path, route.GetTo(), final)
			} else {
				return
			}
		}
	}
}

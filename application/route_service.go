package application

import "fmt"

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

	bs := s.checkPath(possibilities)
	fmt.Println(bs)

	// implementar a regra de negócio para encontrar a melhor rota aqui
	best := BestRoute{FlyPath: "OOASD-ASD-D-ASD-", TotalCost: 90}
	return best, nil
}

func (s *RouteService) checkPath(routes []Route) []BestRoute {
	var bests []BestRoute
	for _, route := range routes {
		b := BestRoute{
			FlyPath: fmt.Sprintf("%s-%s", route.GetFrom(), route.GetTo()),
		}
		c := s.searchInRoutes(route.GetTo())
		if c.From != "" {
			b.FlyPath = fmt.Sprintf("%s-%s", b.GetFlyPath(), c.GetTo())
			b.TotalCost = b.GetTotalCost() + c.GetPrice()
		}

		bests = append(bests, b)

	}

	return bests
}

func (s *RouteService) searchInRoutes(from string) Route {
	routes, _ := s.List()
	var r Route
	for _, route := range routes {
		if route.GetFrom() == from {
			r = route
		}
	}
	return r
}

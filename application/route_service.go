package application

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

func (s *RouteService) List() ([]RouteInterface, error) {
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

	// implementar a regra de negócio para encontrar a melhor rota aqui
	best := BestRoute{FlyPath: "OOASD-ASD-D-ASD-", TotalCost: 90}
	return best, nil
}

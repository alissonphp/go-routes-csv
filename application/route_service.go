package application

type RouteService struct {
	Persistence RoutePersistenceInterface
}

func (s *RouteService) List() ([]RouteInterface, error) {
	routes, err := s.Persistence.List()
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func (s *RouteService) Get(from string, to string) (RouteInterface, error) {
	route, err := s.Persistence.Get(from, to)
	if err != nil {
		return nil, err
	}
	return route, nil

}

func (s *RouteService) Create(from string, to string, price float64) (RouteInterface, error) {
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

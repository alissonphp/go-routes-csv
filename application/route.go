package application

type RouteInterface interface {
	GetFrom() string
	GetTo() string
	GetPrice() float64
}

type RouteServiceInterface interface {
	Create(from string, to string, price float64) (RouteInterface, error)
	SearchBest(from string, to string) (string, error)
}

type RouteReader interface {
	Get(from string, to string) (RouteInterface, error)
}

type RouteWriter interface {
	Save(route RouteInterface) (RouteInterface, error)
}

// RoutePersistenceInterface interface segregation
type RoutePersistenceInterface interface {
	RouteReader
	RouteWriter
}

type Route struct {
	From  string  `json:"from"`
	To    string  `json:"to"`
	Price float64 `json:"price"`
}

func NewRoute() *Route {
	route := Route{}
	return &route
}

func (r *Route) GetFrom() string {
	return r.From
}

func (r *Route) GetTo() string {
	return r.To
}

func (r *Route) GetPrice() float64 {
	return r.Price
}

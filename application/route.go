package application

type RouteInterface interface {
	GetFrom() string
	GetTo() string
	GetPrice() int
}

type RouteServiceInterface interface {
	Save(from string, to string, price int) (RouteInterface, error)
	SearchBest(from string, to string) (BestRoute, error)
}

type RouteReader interface {
	List() ([]Route, error)
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
	From  string `json:"from"`
	To    string `json:"to"`
	Price int    `json:"price"`
}

type BestRoute struct {
	FlyPath   string `json:"fly_path"`
	TotalCost int    `json:"total_cost"`
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

func (r *Route) GetPrice() int {
	return r.Price
}

func (b *BestRoute) GetFlyPath() string {
	return b.FlyPath
}

func (b *BestRoute) GetTotalCost() int {
	return b.TotalCost
}

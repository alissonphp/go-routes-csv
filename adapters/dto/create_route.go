package dto

import (
	"errors"
	"go-best-route/application"
)

type CreateRoute struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Price int    `json:"price"`
}

func NewRoute() *CreateRoute {
	return &CreateRoute{}
}

func (r *CreateRoute) Bind(route *application.Route) (*application.Route, error) {
	if r.From != "" && r.To != "" && r.Price > 0 {
		route.From = r.From
		route.To = r.To
		route.Price = r.Price
	} else {
		return &application.Route{}, errors.New("new route is not valid")
	}

	return route, nil

}

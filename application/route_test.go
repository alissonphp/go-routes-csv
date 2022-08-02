package application_test

import (
	"github.com/stretchr/testify/require"
	"go-best-route/application"
	"testing"
)

func TestRoute_GetFrom(t *testing.T) {
	route := application.Route{}
	route.From = "GRU"
	route.To = "BRC"
	route.Price = 10
	require.Equal(t, route.From, route.GetFrom())
}

func TestRoute_GetTo(t *testing.T) {
	route := application.Route{}
	route.From = "GRU"
	route.To = "BRC"
	route.Price = 10
	require.Equal(t, route.To, route.GetTo())
}

func TestRoute_GetPrice(t *testing.T) {
	route := application.Route{}
	route.From = "GRU"
	route.To = "BRC"
	route.Price = 10
	require.Equal(t, route.Price, route.GetPrice())
}

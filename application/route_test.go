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

func TestBestRoute_GetFlyPath(t *testing.T) {
	bestRoute := application.BestRoute{}
	bestRoute.FlyPath = "GRU-ORL-BSB-FOR"
	bestRoute.TotalCost = 190
	require.Equal(t, bestRoute.FlyPath, bestRoute.GetFlyPath())
}

func TestBestRoute_GetTotalCost(t *testing.T) {
	bestRoute := application.BestRoute{}
	bestRoute.FlyPath = "GRU-ORL-BSB-FOR"
	bestRoute.TotalCost = 190
	require.Equal(t, bestRoute.TotalCost, bestRoute.GetTotalCost())
}

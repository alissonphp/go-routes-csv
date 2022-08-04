package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go-best-route/adapters/cli"
	"go-best-route/application"
	mock_application "go-best-route/application/mocks"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	b := application.BestRoute{
		FlyPath:   "GRU-FOR-STC",
		TotalCost: 30,
	}
	r1 := application.Route{From: "GRU", To: "FOR", Price: 10}
	r2 := application.Route{From: "FOR", To: "STC", Price: 20}

	var routes []application.Route
	routes = append(routes, r1, r2)
	persistence := mock_application.NewMockRoutePersistenceInterface(ctrl)
	persistence.EXPECT().List().Return(routes, nil).AnyTimes()
	service := mock_application.NewMockRouteServiceInterface(ctrl)
	service.EXPECT().SearchBest(gomock.Any(), gomock.Any()).Return(b, nil).AnyTimes()

	res, err := cli.Run(service, "GRU", "STC")
	require.Nil(t, err)
	require.Equal(t, res, fmt.Sprintf("best route: %s > $%d", b.GetFlyPath(), b.GetTotalCost()))
}

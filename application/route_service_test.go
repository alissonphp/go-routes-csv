package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go-best-route/application"
	mock_application "go-best-route/application/mocks"
	"testing"
)

//func TestRouteService_List(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	var routes []application.RouteInterface
//
//	route1 := mock_application.NewMockRouteInterface(ctrl)
//	route2 := mock_application.NewMockRouteInterface(ctrl)
//
//	routes = append(routes, &route1, &route2)
//	persistence := mock_application.NewMockRoutePersistenceInterface(ctrl)
//	persistence.EXPECT().List().Return(routes, nil).AnyTimes()
//
//	service := application.RouteService{
//		Persistence: persistence,
//	}
//
//	res, err := service.List()
//	require.Nil(t, err)
//	require.Equal(t, routes, res)
//}

func TestRouteService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	route := mock_application.NewMockRouteInterface(ctrl)
	persistence := mock_application.NewMockRoutePersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any(), gomock.Any()).Return(route, nil).AnyTimes()

	service := application.RouteService{
		Persistence: persistence,
	}

	res, err := service.Get("any", "any")
	require.Nil(t, err)
	require.Equal(t, route, res)
}

func TestRouteService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	route := mock_application.NewMockRouteInterface(ctrl)
	persistence := mock_application.NewMockRoutePersistenceInterface(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(route, nil).AnyTimes()

	service := application.RouteService{
		Persistence: persistence,
	}

	res, err := service.Save("FOR", "BSB", 10)
	require.Nil(t, err)
	require.Equal(t, route, res)
}

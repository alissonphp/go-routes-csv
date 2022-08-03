package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-best-route/adapters/csv"
	server2 "go-best-route/adapters/web/server"
	"go-best-route/application"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start webserver",
	Run: func(cmd *cobra.Command, args []string) {
		server := server2.MakeNewWebserver()
		routeCsvAdapter := csv.NewReadCsvFile("input-routes.csv")
		server.Service = application.NewRouteService(routeCsvAdapter)
		fmt.Println("The best route http webserver has been started ")
		server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}

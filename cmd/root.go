package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-best-route/adapters/cli"
	"go-best-route/adapters/csv"
	"go-best-route/application"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "myway [csv file] [from-to]",
	Short: "Find the better (and cheapest) way to your destiny",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args[1]) != 7 {
			fmt.Println("error: your route should be in this format: FROM-TO")
			os.Exit(1)
		}

		path := strings.Split(args[1], "-")
		routeCsvAdapter := csv.NewReadCsvFile(args[0])
		routeService := application.NewRouteService(routeCsvAdapter)
		res, _ := cli.Run(routeService, path[0], path[1])
		fmt.Println(res)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

package main

import (
	"encoding/csv"
	"fmt"
	"go-best-route/application"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.OpenFile("input-routes.csv", os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	if err := w.Write([]string{"OOO", "LOL", "12"}); err != nil {
		log.Fatal(err)
	}

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		s, err := strconv.ParseFloat(line[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		route := application.Route{
			From:  line[0],
			To:    line[1],
			Price: s,
		}
		fmt.Println(route)
	}
}

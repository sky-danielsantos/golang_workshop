package main

import (
	"fmt"
	"os"
)

func getData() ([]string, error) {
	return []string{"There", "Are", "No", "Pitfalls", "In", "Golang"}, nil
}

func assignPitfall() {
	var data []string

	abortProcessData := os.Getenv("ABORT_PROCESS")

	if abortProcessData == "" {
		fmt.Println("Processing Data")
		data, err := getData()
		if err != nil {
			panic("ERROR fetching data")
		}

		fmt.Printf("Data was fetched %d\n", len(data))
	}

	for _, item := range data {
		fmt.Println(item)
	}
}

func main() {
	assignPitfall()
}

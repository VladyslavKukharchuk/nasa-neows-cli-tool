package main

import (
	"fmt"
	"log"
	"nasa-neows-cli-tool/neows"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	const URL = "https://api.nasa.gov/neo/rest/v1/"
	apiKey := os.Getenv("API_KEY")

	neoWsJSON, err := neows.GetNEOsByDaysAgo(URL, apiKey, 7)
	if err != nil {
		return err
	}

	fmt.Println(neoWsJSON)

	return nil
}

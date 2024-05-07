package main

import (
	"fmt"
	"log"
	"nasa-neows-cli-tool/neows"
)

func main() {
	defer handlerPanic()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	neoWsJSON, err := neows.GetNEOsByDaysAgo(7)
	if err != nil {
		return err
	}

	fmt.Println(neoWsJSON)

	return nil
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

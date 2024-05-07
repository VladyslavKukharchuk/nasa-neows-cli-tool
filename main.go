package main

import (
	"fmt"
	"nasa-neows-cli-tool/NeoWs"
)

func main() {
	defer handlerPanic()

	neoWsJSON := NeoWs.GetNEOsByDaysAgo(7)

	fmt.Println(neoWsJSON)
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

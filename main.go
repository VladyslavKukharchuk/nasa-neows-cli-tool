package main

import (
	"fmt"
	"nasa-neows-cli-tool/NeoWs"
)

func main() {
	defer handlerPanic()

	neoWsJson := NeoWs.ShowNEOsByLastSevenDays()

	fmt.Println(neoWsJson)
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

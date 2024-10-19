package main

import (
	"fmt"
	"strings"

	kiteinstruments "github.com/nsvirk/gokiteinstruments"
)

func printResults(title string, instruments []kiteinstruments.Instrument, limit int) {
	sLine := strings.Repeat("-", 60)
	fmt.Println(sLine)
	fmt.Printf("%s => %d instruments\n", title, len(instruments))
	fmt.Println(sLine)

	if limit > 0 && limit < len(instruments) {
		instruments = instruments[:limit]
	}

	for i, instrument := range instruments {
		fmt.Printf("%d. %s:%s\n", i+1, instrument.Exchange, instrument.Tradingsymbol)
	}
	fmt.Println()
}

func printMapResults(title string, results map[string][]string) {
	sLine := strings.Repeat("-", 60)
	fmt.Println(sLine)
	fmt.Printf("%s => %d segments\n", title, len(results))
	fmt.Println(sLine)

	for key, values := range results {
		fmt.Printf("%s: %v\n", key, values)
	}
	fmt.Println()
}

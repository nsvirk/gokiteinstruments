package main

import (
	"context"
	"fmt"
	"log"
	"time"

	kiteinstruments "github.com/nsvirk/gokiteinstruments"
)

func main() {
	ctx := context.Background()
	client, err := kiteinstruments.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	sLine := "-------------------------------------------------------------"

	// Query by instrument_token 408065
	instrumentToken := uint32(408065)
	instruments := client.QueryByInstrumentToken(instrumentToken)
	fmt.Println(sLine)
	fmt.Printf("QueryByInstrumentToken(%d) => %d instruments\n", instrumentToken, len(instruments))
	fmt.Println(sLine)
	for i, instrument := range instruments {
		fmt.Printf("%d. %s:%s\n", i+1, instrument.Exchange, instrument.Tradingsymbol)
	}
	fmt.Println("")

	// Query by tradingsymbol
	tradingsymbol := "SBIN"
	instruments = client.QueryByTradingsymbol(tradingsymbol)
	fmt.Println(sLine)
	fmt.Printf("QueryByTradingsymbol(%s) => %d instruments\n", tradingsymbol, len(instruments))
	fmt.Println(sLine)
	for i, instrument := range instruments {
		fmt.Printf("%d. %s:%s\n", i+1, instrument.Exchange, instrument.Tradingsymbol)
	}
	fmt.Println("")

	// Query by segment
	segment := "MCX-FUT"
	instruments = client.QueryBySegment(segment)
	fmt.Println(sLine)
	fmt.Printf("QueryBySegment(%s) => %d instruments\n", segment, len(instruments))
	fmt.Println(sLine)
	for i, instrument := range instruments[:5] {
		fmt.Printf("%d. %s:%s\n", i+1, instrument.Exchange, instrument.Tradingsymbol)
	}
	fmt.Println("")

	// Query by name
	name := "NIFTY"
	instruments = client.QueryByName(name)
	fmt.Println(sLine)
	fmt.Printf("QueryByName(%s) => %d instruments\n", name, len(instruments))
	fmt.Println(sLine)
	for i, instrument := range instruments[:5] {
		fmt.Printf("%d. %s:%s\n", i+1, instrument.Exchange, instrument.Tradingsymbol)
	}

	// Query by exchange and trading symbol
	exchange := "NSE"
	tradingsymbol = "SBIN"
	instruments = client.QueryByExchangeTradingSymbol(exchange, tradingsymbol)
	fmt.Println(sLine)
	fmt.Printf("QueryByExchangeTradingSymbol(%s, %s) => %d instruments\n", exchange, tradingsymbol, len(instruments))
	fmt.Println(sLine)
	for i, instrument := range instruments {
		fmt.Printf("%d. %s:%s\n", i+1, instrument.Exchange, instrument.Tradingsymbol)
	}
	fmt.Println("")

	// Query by segment and name
	segment = "NFO-FUT"
	name = "FINNIFTY"
	instruments = client.QueryBySegmentName(segment, name)
	fmt.Println(sLine)
	fmt.Printf("QueryBySegmentName(%s, %s) => %d instruments\n", segment, name, len(instruments))
	fmt.Println(sLine)
	for i, instrument := range instruments {
		fmt.Printf("%d. %s:%s\n", i+1, instrument.Exchange, instrument.Tradingsymbol)
	}
	fmt.Println("")

	// Query segment expiries by names
	segmentExpiries := client.QuerySegmentExpiriesByName(name)
	fmt.Println(sLine)
	fmt.Printf("QuerySegmentExpiriesByName(%s) => %d segments\n", name, len(segmentExpiries))
	fmt.Println(sLine)
	for segment, expiries := range segmentExpiries {
		fmt.Printf("%s: %v\n", segment, expiries)
	}
	fmt.Println("")

	// Query segment names by expiry
	expiry := time.Now().Format("2006-01-02")
	segmentNames := client.QuerySegmentNamesByExpiry(expiry)
	fmt.Println(sLine)
	fmt.Printf("QuerySegmentNamesByExpiry(%s) => %d segments\n", expiry, len(segmentNames))
	fmt.Println(sLine)
	for segment, names := range segmentNames {
		fmt.Printf("%s: %v\n", segment, names)
	}
	fmt.Println("")

	fmt.Println(sLine)
}

package main

import (
	"fmt"
	"time"

	kiteinstruments "github.com/nsvirk/gokiteinstruments"
)

func runQueries(client *kiteinstruments.Client) {
	// Query by instrument_token
	instrumentToken := uint32(408065)
	instruments := client.QueryByInstrumentToken(instrumentToken)
	printResults(fmt.Sprintf("QueryByInstrumentToken(%d)", instrumentToken), instruments, 0)

	// Query by tradingsymbol
	tradingsymbol := "SBIN"
	instruments = client.QueryByTradingsymbol(tradingsymbol)
	printResults(fmt.Sprintf("QueryByTradingsymbol(%s)", tradingsymbol), instruments, 0)

	// Query by segment
	segment := "MCX-FUT"
	instruments = client.QueryBySegment(segment)
	printResults(fmt.Sprintf("QueryBySegment(%s)", segment), instruments, 5)

	// Query by name
	name := "NIFTY"
	instruments = client.QueryByName(name)
	printResults(fmt.Sprintf("QueryByName(%s)", name), instruments, 5)

	// Query by exchange and trading symbol
	exchange := "NSE"
	tradingsymbol = "SBIN"
	instruments = client.QueryByExchangeTradingSymbol(exchange, tradingsymbol)
	printResults(fmt.Sprintf("QueryByExchangeTradingSymbol(%s, %s)", exchange, tradingsymbol), instruments, 0)

	// Query by segment and name
	segment = "NFO-FUT"
	name = "FINNIFTY"
	instruments = client.QueryBySegmentName(segment, name)
	printResults(fmt.Sprintf("QueryBySegmentName(%s, %s)", segment, name), instruments, 0)

	// Query segment expiries by names
	segmentExpiries := client.QuerySegmentExpiriesByName(name)
	printMapResults(fmt.Sprintf("QuerySegmentExpiriesByName(%s)", name), segmentExpiries)

	// Query segment names by expiry
	expiry := time.Now().Format("2006-01-02")
	segmentNames := client.QuerySegmentNamesByExpiry(expiry)
	printMapResults(fmt.Sprintf("QuerySegmentNamesByExpiry(%s)", expiry), segmentNames)
}

package kiteinstruments

import (
	"context"
	"sort"
)

const InstrumentsURL = "https://api.kite.trade/instruments"

// Client represents a client for the kiteinstruments library
type Client struct {
	instruments []Instrument
}

// NewClient creates a new Client and fetches the instruments data
func NewClient(ctx context.Context) (*Client, error) {
	instruments, err := FetchInstruments(ctx, InstrumentsURL)
	if err != nil {
		return nil, err
	}
	return &Client{instruments: instruments}, nil
}

// Query allows querying the instruments data
func (c *Client) query(filter func(Instrument) bool) []Instrument {
	return queryInstruments(c.instruments, filter)
}

// QueryByInstrumentToken allows querying the instruments data by instrument token
func (c *Client) QueryByInstrumentToken(instrumentToken uint32) []Instrument {
	return c.query(func(i Instrument) bool {
		return i.InstrumentToken == instrumentToken
	})
}

// QueryByTradingsymbol allows querying the instruments data by trading symbol
func (c *Client) QueryByTradingsymbol(tradingsymbol string) []Instrument {
	return c.query(func(i Instrument) bool {
		return i.Tradingsymbol == tradingsymbol
	})
}

// QueryBySegment allows querying the instruments data by segment
func (c *Client) QueryBySegment(segment string) []Instrument {
	return c.query(func(i Instrument) bool {
		return i.Segment == segment
	})
}

// QueryByName allows querying the instruments data by name
func (c *Client) QueryByName(name string) []Instrument {
	return c.query(func(i Instrument) bool {
		return i.Name == name
	})
}

// QueryByExpiry allows querying the instruments data by expiry
func (c *Client) QueryByExpiry(expiry string) []Instrument {
	return c.query(func(i Instrument) bool {
		return i.Expiry == expiry
	})
}

// QueryByExchangeTradingSymbol allows querying the instruments data by exchange and trading symbol
func (c *Client) QueryByExchangeTradingSymbol(exchange, tradingsymbol string) []Instrument {
	return c.query(func(i Instrument) bool {
		return i.Exchange == exchange && i.Tradingsymbol == tradingsymbol
	})
}

// QueryBySegmentName allows querying the instruments data by segment name
func (c *Client) QueryBySegmentName(segment, name string) []Instrument {
	return c.query(func(i Instrument) bool {
		return i.Segment == segment && i.Name == name
	})
}

// QuerySegmentExpiriesByName allows querying the instruments data by name
func (c *Client) QuerySegmentExpiriesByName(name string) map[string][]string {
	instrument := c.QueryByName(name)
	segmentMap := make(map[string]map[string]map[string]bool)
	for _, instrument := range instrument {
		nameKey := instrument.Name
		segmentKey := instrument.Segment
		expiryKey := instrument.Expiry

		if _, exists := segmentMap[nameKey]; !exists {
			segmentMap[nameKey] = make(map[string]map[string]bool)
		}
		if _, exists := segmentMap[nameKey][segmentKey]; !exists {
			segmentMap[nameKey][segmentKey] = make(map[string]bool)
		}
		segmentMap[nameKey][segmentKey][expiryKey] = true
	}

	// Convert the map of sets to a map of slices
	segmentExpiries := make(map[string][]string)
	for _, segmentSet := range segmentMap {
		for segmentKey, expirySet := range segmentSet {
			expiries := make([]string, 0, len(expirySet))
			for expiry := range expirySet {
				expiries = append(expiries, expiry)
			}
			sort.Strings(expiries)
			segmentExpiries[segmentKey] = expiries
		}
	}
	return segmentExpiries
}

// QuerySegmentNamesByExpiry allows querying the instruments data by expiry
func (c *Client) QuerySegmentNamesByExpiry(expiry string) map[string][]string {
	instruments := c.QueryByExpiry(expiry)
	expiryMap := make(map[string]map[string]map[string]bool)
	for _, instrument := range instruments {
		nameKey := instrument.Name
		segmentKey := instrument.Segment
		expiryKey := instrument.Expiry

		if _, exists := expiryMap[expiryKey]; !exists {
			expiryMap[expiryKey] = make(map[string]map[string]bool)
		}
		if _, exists := expiryMap[expiryKey][segmentKey]; !exists {
			expiryMap[expiryKey][segmentKey] = make(map[string]bool)
		}
		expiryMap[expiryKey][segmentKey][nameKey] = true
	}
	// Convert the map of sets to a map of slices
	segmentNames := make(map[string][]string)
	for _, segmentSet := range expiryMap {
		for segmentKey, nameSet := range segmentSet {
			names := make([]string, 0, len(nameSet))
			for name := range nameSet {
				names = append(names, name)
			}
			sort.Strings(names)
			segmentNames[segmentKey] = names
		}
	}
	return segmentNames
}

package kiteinstruments

// Instrument represents a financial instrument
type Instrument struct {
	InstrumentToken uint32  `csv:"instrument_token"`
	ExchangeToken   uint32  `csv:"exchange_token"`
	Tradingsymbol   string  `csv:"tradingsymbol"`
	Name            string  `csv:"name"`
	LastPrice       float64 `csv:"last_price"`
	Expiry          string  `csv:"expiry"`
	StrikePrice     float64 `csv:"strike"`
	TickSize        float64 `csv:"tick_size"`
	LotSize         int     `csv:"lot_size"`
	InstrumentType  string  `csv:"instrument_type"`
	Segment         string  `csv:"segment"`
	Exchange        string  `csv:"exchange"`
}

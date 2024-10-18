package kiteinstruments

// queryInstruments filters the instruments based on the provided filter function
func queryInstruments(instruments []Instrument, filter func(Instrument) bool) []Instrument {
	var result []Instrument
	for _, instrument := range instruments {
		if filter(instrument) {
			result = append(result, instrument)
		}
	}
	return result
}

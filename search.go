package instruments

// GetByID returns an instrument using the symbol.
func (m *Manager) GetByID(id string) (Instrument, error) {
	inst, ok := m.idToInst[id]
	if !ok {
		return Instrument{}, ErrInstrumentNotFound
	}
	return *inst, nil
}

// GetByExchTradingsymbol returns an instrument using exchange and trading symbol.
func (m *Manager) GetByExchTradingsymbol(exchange, tradingsymbol string) (Instrument, error) {
	return m.GetByID(exchange + ":" + tradingsymbol)
}

// GetByISIN returns a set of instruments using ISIN.
func (m *Manager) GetByISIN(isin string) ([]Instrument, error) {
	insts, ok := m.isinToInstruments[isin]
	if !ok {
		return []Instrument{}, ErrInstrumentNotFound
	}

	out := make([]Instrument, 0, len(insts))
	for _, i := range insts {
		out = append(out, *i)
	}

	return out, nil
}

// GetByInstToken returns an instrument using instrument token.
func (m *Manager) GetByInstToken(token uint32) (Instrument, error) {
	inst, ok := m.tokenToInstrument[token]
	if !ok {
		return Instrument{}, ErrInstrumentNotFound
	}
	return *inst, nil
}

// GetByExchToken takes an exchange token.
func (m *Manager) GetByExchToken(exch string, exchToken uint32) (Instrument, error) {
	// Get the segment ID from the exchange.
	segID, found := m.segmentIDs[exch]
	if !found {
		return Instrument{}, ErrSegmentNotFound
	}

	instToken := ExchTokenToInstToken(segID, exchToken)
	inst, found := m.tokenToInstrument[instToken]
	if !found {
		return Instrument{}, ErrInstrumentNotFound
	}
	return *inst, nil
}

// Filter returns a list of instruments filtered by the given filter.
func (m *Manager) Filter(filter func(Instrument) bool) []Instrument {
	out := []Instrument{}

	for _, v := range m.tokenToInstrument {
		if filter(*v) {
			out = append(out, *v)
		}
	}

	return out
}

// GetAllByUnderlying returns a list of F&O instruments associated with the underlying tradingsymbol.
func (m *Manager) GetAllByUnderlying(exchange, underlying string) ([]Instrument, error) {
	out := []Instrument{}

	for _, ins := range m.tokenToInstrument {
		if ins.Exchange == exchange && ins.Name == underlying {
			out = append(out, *ins)
		}
	}

	if len(out) == 0 {
		return []Instrument{}, ErrInstrumentNotFound
	}

	return out, nil
}

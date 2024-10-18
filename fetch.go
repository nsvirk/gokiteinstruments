package kiteinstruments

import (
	"context"
	"net/http"

	"github.com/gocarina/gocsv"
)

// FetchInstruments fetches and parses the CSV data from the given URL
func FetchInstruments(ctx context.Context, url string) ([]Instrument, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var instruments []Instrument
	if err := gocsv.Unmarshal(resp.Body, &instruments); err != nil {
		return nil, err
	}

	return instruments, nil
}

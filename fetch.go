package kiteinstruments

import (
	"context"
	"net/http"
	"time"

	"github.com/gocarina/gocsv"
)

const requestTimeout = 7 * time.Second

// fetchInstruments fetches and parses the CSV data from the given URL
func fetchInstruments(ctx context.Context, url string) ([]Instrument, error) {
	client := &http.Client{
		Timeout: requestTimeout,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, ErrInvalidURL
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, ErrFetchFailed
	}
	defer resp.Body.Close()

	var instruments []Instrument
	if err := gocsv.Unmarshal(resp.Body, &instruments); err != nil {
		return nil, ErrParsingFailed
	}

	return instruments, nil
}

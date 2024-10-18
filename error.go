package kiteinstruments

import "errors"

var (
	ErrInvalidURL    = errors.New("invalid URL provided")
	ErrFetchFailed   = errors.New("failed to fetch instruments data")
	ErrParsingFailed = errors.New("failed to parse instruments data")
)

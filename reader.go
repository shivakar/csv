package csv

import (
	lcsv "encoding/csv"
	"io"
)

// Reader is a drop-in replacement for encoding/csv.Reader
type Reader struct {
	*lcsv.Reader
}

// NewReader returns a new CSV reader
func NewReader(r io.Reader) *Reader {
	return &Reader{
		lcsv.NewReader(r),
	}
}

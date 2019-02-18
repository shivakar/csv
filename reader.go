package csv

import (
	lcsv "encoding/csv"
	"io"
)

// Reader is a drop-in replacement for encoding/csv.Reader
type Reader struct {
	reader *lcsv.Reader
}

// NewReader returns a new CSV reader
func NewReader(r io.Reader) *Reader {
	rdr := &Reader{}
	lr := lcsv.NewReader(r)
	rdr.reader = lr
	return rdr
}

// Read reads the next record from underlying reader
func (r *Reader) Read() ([]string, error) {
	return r.reader.Read()
}

// ReadAll reads all remaining records from underlying reader
func (r *Reader) ReadAll() ([][]string, error) {
	return r.reader.ReadAll()
}

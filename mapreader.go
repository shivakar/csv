package csv

import (
	lcsv "encoding/csv"
	"errors"
	"fmt"
	"io"
)

var (
	ErrReadHeader = errors.New("error reading header")
)

// MapReader represents a CSV reader
type MapReader struct {
	Header []string
	reader *lcsv.Reader
}

// NewMapReader returns a CSV reader that represents a row as a map of column names to values
func NewMapReader(r io.Reader) (*MapReader, error) {
	m := &MapReader{}
	m.reader = lcsv.NewReader(r)
	cols, err := m.reader.Read()
	if err != nil {
		return nil, fmt.Errorf("%v: %v", ErrReadHeader, err)
	}
	m.Header = cols
	return m, nil
}

// Read reads the next record in CSV and returns a map of header columns to row values
func (m *MapReader) Read() (map[string]string, error) {
	row, err := m.reader.Read()
	if err != nil {
		return nil, err
	}

	output := map[string]string{}
	for i, val := range row {
		output[m.Header[i]] = val
	}
	return output, nil
}

// ReadAll reads all the record in CSV and returns a slice of map[string]string
func (m *MapReader) ReadAll() ([]map[string]string, error) {
	rows, err := m.reader.ReadAll()
	if err != nil {
		return nil, err
	}

	output := []map[string]string{}
	for _, row := range rows {
		rec := map[string]string{}
		for i, val := range row {
			rec[m.Header[i]] = val
		}
		output = append(output, rec)
	}
	return output, nil
}

package csv

import (
	lcsv "encoding/csv"
	"errors"
	"reflect"
	"strings"
	"testing"
)

type TestCase struct {
	Label          string
	Input          string
	Expected       []map[string]string
	NewReaderError error
	ReadError      error
}

func getTestCases() []TestCase {
	return []TestCase{
		{
			Label: "simple",
			Input: "header1,header2,header3\nrow1col1,row1col2,row1col3\nrow2col1,row2col2,row2col3",
			Expected: []map[string]string{
				{"header1": "row1col1", "header2": "row1col2", "header3": "row1col3"},
				{"header1": "row2col1", "header2": "row2col2", "header3": "row2col3"},
			},
			NewReaderError: nil,
			ReadError:      nil,
		},
		{
			Label:          "nocontent",
			Input:          "",
			Expected:       nil,
			NewReaderError: errors.New("error reading header: EOF"),
			ReadError:      nil,
		},
		{
			Label:          "invalid CSV",
			Input:          "header1,header2\nval1,val2,val3",
			Expected:       nil,
			NewReaderError: nil,
			ReadError:      &lcsv.ParseError{StartLine: 2, Line: 2, Column: 0, Err: lcsv.ErrFieldCount},
		},
	}
}

func cmpNewReaderError(tt *TestCase, err error, t *testing.T) {
	if !reflect.DeepEqual(err, tt.NewReaderError) {
		t.Errorf("NewMapReader(): unexpected error:\nExpected: %v, got: %v\nTest case: %s", tt.NewReaderError, err, tt.Label)
	}
}

func TestMapReaderRead(t *testing.T) {
	tests := getTestCases()
	for _, tt := range tests {
		r, err := NewMapReader(strings.NewReader(tt.Input))
		cmpNewReaderError(&tt, err, t)
		if err != nil {
			continue
		}
		for i := 0; i == 0 || i < len(tt.Expected); i++ {
			actual, err := r.Read()
			if err != nil {
				if !reflect.DeepEqual(err, tt.ReadError) {
					t.Errorf("Read() error:\nExpected: %v, got: %v\nTest case: %s", tt.ReadError, err, tt.Label)
					break
				}
			}
			if i >= len(tt.Expected) {
				break
			}
			expected := tt.Expected[i]
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Read output:\nExpected: %q\n, got: %q\nTest case: %s", expected, actual, tt.Label)
			}
		}
	}
}

func TestMapReaderReadAll(t *testing.T) {
	tests := getTestCases()
	for _, tt := range tests {
		r, err := NewMapReader(strings.NewReader(tt.Input))
		cmpNewReaderError(&tt, err, t)
		if err != nil {
			continue
		}
		out, err := r.ReadAll()
		if !reflect.DeepEqual(err, tt.ReadError) {
			t.Errorf("ReadAll() error:\nExpected: %v, got: %v\nTest case: %s", tt.ReadError, err, tt.Label)
			continue
		}
		if !reflect.DeepEqual(out, tt.Expected) {
			t.Errorf("ReadAll output:\nExpected: %q\n, got: %q\nTest case: %s", tt.Expected, out, tt.Label)
		}
	}
}

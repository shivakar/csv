package csv

import (
	"reflect"
	"strings"
	"testing"
)

func TestReaderRead(t *testing.T) {
	tests := []struct {
		Input    string
		Expected [][]string
		Error    error
	}{
		{
			Input:    "a,b,c\nd,e,f",
			Expected: [][]string{{"a", "b", "c"}, {"d", "e", "f"}},
			Error:    nil,
		},
	}

	for _, tt := range tests {
		r := NewReader(strings.NewReader(tt.Input))
		for _, expected := range tt.Expected {
			actual, err := r.Read()
			if err != nil {
				if !reflect.DeepEqual(err, tt.Error) {
					t.Errorf("Read() error:\nExpected: %v, got %v", tt.Error, err)
				}
			}
			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("Read() output:\n Expected: %v, got %v", expected, actual)
			}
		}
	}
}

func TestReaderReadAll(t *testing.T) {
	tests := []struct {
		Input    string
		Expected [][]string
		Error    error
	}{
		{
			Input:    "a,b,c\n",
			Expected: [][]string{{"a", "b", "c"}},
			Error:    nil,
		},
	}

	for _, tt := range tests {
		r := NewReader(strings.NewReader(tt.Input))
		out, err := r.ReadAll()
		if !reflect.DeepEqual(err, tt.Error) {
			t.Errorf("ReadAll() error:\nExpected: %v, got: %v", tt.Error, err)
			continue
		}
		if !reflect.DeepEqual(out, tt.Expected) {
			t.Errorf("ReadAll output:\nExpected: %q\n, got: %q", tt.Expected, out)
		}
	}
}

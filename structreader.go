package csv

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	// CSVTag is the tag name used for mapping fields to CSV header names
	CSVTag = "csv"
)

// StructReader returns a CSV reader that can read records into structs
type StructReader struct {
	m *MapReader
}

// NewStructReader returns a new StructReader
func NewStructReader(r io.Reader) (*StructReader, error) {
	s := &StructReader{}
	sm, err := NewMapReader(r)
	if err != nil {
		return nil, err
	}
	s.m = sm
	return s, nil
}

// Read reads the next record into the provided interface
func (s *StructReader) Read(v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("read error: non-pointer type %v", rv.Type())
	}
	if rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("read error: non-struct kind %v", rv.Elem().Kind())
	}

	mv, err := s.m.Read()
	if err != nil {
		return err
	}

	elem := rv.Elem()
	for i := 0; i < elem.NumField(); i++ {
		typ := elem.Type().Field(i)
		col := typ.Tag.Get(CSVTag)
		if col == "-" {
			continue
		}
		val, ok := mv[col]
		if !ok {
			return fmt.Errorf("unknown column: %s", col)
		}

		field := elem.Field(i)
		switch field.Type().Name() {
		case "string":
			field.SetString(val)
		case "bool":
			switch strings.ToLower(val) {
			case "true":
				field.SetBool(true)
			case "false":
				field.SetBool(false)
			default:
				return fmt.Errorf("invalid value for bool: %s", val)
			}
		case "int":
			iv, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(iv)
		case "float64":
			fv, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			field.SetFloat(fv)
		case "Time":
			fv, err := time.Parse("2006-01-02 15:04:05 -0700 MST", val)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(fv))
		default:
			return fmt.Errorf("unknown field type: %v", field.Type().Name())
		}
	}

	return nil
}

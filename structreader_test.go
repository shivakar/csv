package csv

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

type T1 struct {
	S      string    `csv:"string"`
	B      bool      `csv:"bool"`
	I      int       `csv:"int"`
	F      float64   `csv:"float64"`
	Ignore string    `csv:"-"`
	S2     string    `csv:"s2"`
	T      time.Time `csv:"time"`
}

func TestStructReaderRead(t *testing.T) {
	testdata := "string,bool,int,float64,s2,time\nhello world,false,42,3.14,tardis,2019-02-22 00:07:37 -0500 EST"
	ts, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", "2019-02-22 00:07:37 -0500 EST")
	expected := T1{
		S:      "hello world",
		B:      false,
		I:      42,
		F:      3.14,
		Ignore: "",
		S2:     "tardis",
		T:      ts,
	}
	r, err := NewStructReader(strings.NewReader(testdata))
	if err != nil {
		t.Errorf("NewStructReader() unexpected error: %v", err)
	}

	t1 := T1{}
	err = r.Read(&t1)
	if err != nil {
		t.Errorf("Read() unexpected error: %v", err)
	}

	if !reflect.DeepEqual(t1, expected) {
		t.Errorf("Output error:\nExpected: %v, got: %v", expected, t1)
	}
}

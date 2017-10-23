package fedach

import (
	"reflect"
	"testing"
)

func TestUnmarshalToStruct(t *testing.T) {
	cases := []struct {
		input          []byte
		ExpectedOutput []RoutingDirectoryRecord
		ExpectedErr    error
	}{
		{[]byte(`011000015O0110000150122415000000000FEDERAL RESERVE BANK                1000 PEACHTREE ST N.E.              ATLANTA             GA303094470877372245711     `), []RoutingDirectoryRecord{{"011000015", "O", "011000015", "0", "122415", "000000000", "FEDERAL RESERVE BANK                ", "1000 PEACHTREE ST N.E.              ", "ATLANTA             ", "GA", "30309", "4470", "877", "372", "2457", "1", "1", "     "}}, nil},
		{[]byte(`011000015O0110000150122415000000000FEDERAL RESERVE BANK                1000 PEACHTREE ST N.E.              ATLANTA             GA303094470877372245711     
`), []RoutingDirectoryRecord{{"011000015", "O", "011000015", "0", "122415", "000000000", "FEDERAL RESERVE BANK                ", "1000 PEACHTREE ST N.E.              ", "ATLANTA             ", "GA", "30309", "4470", "877", "372", "2457", "1", "1", "     "}}, nil},
	}
	for _, tc := range cases {
		var output []RoutingDirectoryRecord

		err := Unmarshal(tc.input, &output)
		if err != tc.ExpectedErr {
			t.Errorf("Unmarshal func returned wrong error: got %#v want %#v",
				err, tc.ExpectedErr)
		}

		if !reflect.DeepEqual(tc.ExpectedOutput, output) {
			t.Errorf("Unmarshal func returned wrong output: got %#v want %#v",
				output, tc.ExpectedOutput)
		}
	}
}

func TestUnmarshalToSlice(t *testing.T) {
	cases := []struct {
		input          []byte
		ExpectedOutput [][]string
		ExpectedErr    error
	}{
		{[]byte(`011000015O0110000150122415000000000FEDERAL RESERVE BANK                1000 PEACHTREE ST N.E.              ATLANTA             GA303094470877372245711     `), [][]string{{"011000015", "O", "011000015", "0", "122415", "000000000", "FEDERAL RESERVE BANK                ", "1000 PEACHTREE ST N.E.              ", "ATLANTA             ", "GA", "30309", "4470", "877", "372", "2457", "1", "1", "     "}}, nil},
		{[]byte(`011000015O0110000150122415000000000FEDERAL RESERVE BANK                1000 PEACHTREE ST N.E.              ATLANTA             GA303094470877372245711     
`), [][]string{{"011000015", "O", "011000015", "0", "122415", "000000000", "FEDERAL RESERVE BANK                ", "1000 PEACHTREE ST N.E.              ", "ATLANTA             ", "GA", "30309", "4470", "877", "372", "2457", "1", "1", "     "}}, nil},
	}
	for _, tc := range cases {
		var output [][]string

		err := Unmarshal(tc.input, &output)
		if err != tc.ExpectedErr {
			t.Errorf("Unmarshal func returned wrong error: got %#v want %#v",
				err, tc.ExpectedErr)
		}

		if !reflect.DeepEqual(tc.ExpectedOutput, output) {
			t.Errorf("Unmarshal func returned wrong output: got %#v want %#v",
				output, tc.ExpectedOutput)
		}
	}
}

func TestMarshal(t *testing.T) {
	cases := []struct {
		input          []RoutingDirectoryRecord
		ExpectedOutput []byte
		ExpectedErr    error
	}{
		{[]RoutingDirectoryRecord{{"011000015", "O", "011000015", "0", "122415", "000000000", "FEDERAL RESERVE BANK                ", "1000 PEACHTREE ST N.E.              ", "ATLANTA             ", "GA", "30309", "4470", "877", "372", "2457", "1", "1", "     "}}, []byte(`011000015O0110000150122415000000000FEDERAL RESERVE BANK                1000 PEACHTREE ST N.E.              ATLANTA             GA303094470877372245711     `), nil},
	}
	for _, tc := range cases {
		output, err := Marshal(tc.input)
		if err != tc.ExpectedErr {
			t.Errorf("Marshal func returned wrong error: got %#v want %#v",
				err, tc.ExpectedErr)
		}

		if !reflect.DeepEqual(tc.ExpectedOutput, output) {
			t.Errorf("Unmarshal func returned wrong output: got %s want %s\n, output length: %d ExpectedOutput length: %d",
				output, tc.ExpectedOutput, len(output), len(tc.ExpectedOutput))
		}
	}
}

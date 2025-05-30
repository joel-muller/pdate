package internal

import (
	"errors"
	"reflect"
	"testing"
)

func TestSortArguments(t *testing.T) {
	type parseIgnoreTest struct {
		name   string
		args   []string
		output map[option][]string
		err    error
	}

	var parseIgnoreTests = []parseIgnoreTest{
		{
			name: "dates with ignore and reverse",
			args: []string{"2005-10-01", "2003-04-02", "-i", "mo", "tue", "sa", "-r"},
			output: map[option][]string{
				DATE:    {"2005-10-01", "2003-04-02"},
				IGNORE:  {"mo", "tue", "sa"},
				REVERSE: {},
			},
			err: nil,
		},
		{
			name: "only dates",
			args: []string{"2020-01-01", "2021-01-01"},
			output: map[option][]string{
				DATE: {"2020-01-01", "2021-01-01"},
			},
			err: nil,
		},
		{
			name: "ignore option without values",
			args: []string{"2022-01-01", "-i"},
			output: map[option][]string{
				DATE:   {"2022-01-01"},
				IGNORE: {},
			},
			err: nil,
		},
		{
			name: "reverse and ignore mix",
			args: []string{"-r", "z", "y", "-i", "x"},
			output: map[option][]string{
				DATE:    {},
				REVERSE: {"z", "y"},
				IGNORE:  {"x"},
			},
			err: nil,
		},
		{
			name:   "duplicate options",
			args:   []string{"a", "-i", "b", "-r", "c", "-i", "d"},
			output: map[option][]string{},
			err:    errors.New("found duplicate option argument"),
		},
		{
			name: "unknown option treated as date",
			args: []string{"-x", "some", "value"},
			output: map[option][]string{
				DATE: {"-x", "some", "value"},
			},
			err: nil,
		},
		{
			name: "no arguments after options",
			args: []string{"-r", "-i"},
			output: map[option][]string{
				DATE:    {},
				IGNORE:  {},
				REVERSE: {},
			},
			err: nil,
		},
		{
			name: "found unknown argument",
			args: []string{"-a", "-i"},
			output: map[option][]string{
				DATE:    {},
				IGNORE:  {},
				REVERSE: {},
			},
			err: errors.New("found unknown argument"),
		},
		{
			name: "empty argument",
			args: []string{"-i", ""},
			output: map[option][]string{
				DATE:    {},
				IGNORE:  {},
				REVERSE: {},
			},
			err: errors.New("empty argument provided"),
		},
	}
	for _, test := range parseIgnoreTests {
		t.Run(test.name, func(t *testing.T) {
			parsed, err := SortArguments(test.args)
			if test.err != nil {
				if err == nil {
					t.Fatalf("expected an error got nothing: expected: %v and got: %v \n", err, test.err)
				}
				if err.Error() != test.err.Error() {
					t.Fatalf("expected error: %v, got: %v", test.err, err)
				}
				return
			}
			for key, value := range parsed {
				if !reflect.DeepEqual(value, test.output[key]) {
					t.Fatalf("the slices for %v didn't match: %v and %v \n", key, value, test.output[key])
				}
			}
		})
	}
}

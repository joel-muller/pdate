package internal

import (
	"errors"
	"testing"
	"time"
)

func TestGetDates(t *testing.T) {
	type datesTest struct {
		name        string
		args        []string
		lengthDates int
		err         error
	}
	var tests = []datesTest{
		{
			name:        "all normal",
			args:        []string{"2020-04-03"},
			lengthDates: -1,
			err:         nil,
		},
		{
			name:        "all normal, two arrays",
			args:        []string{"2020-04-03", "2020-05-03", "-i", "mo", "tu", "-r"},
			lengthDates: 23,
			err:         nil,
		},
		{
			name:        "invalid date format",
			args:        []string{"invalid-date"},
			lengthDates: 0,
			err:         errors.New("error while parsing dates first date provided is not valid"),
		},
		{
			name:        "invalid ignore argument",
			args:        []string{"2020-04-03", "-i", "invalid-day"},
			lengthDates: 0,
			err:         errors.New("error while parsing the ignored data error, could not parse the weekdays arguments correctly"),
		},
		{
			name:        "invalid reverse argument",
			args:        []string{"2020-04-03", "-r", "unexpected-arg"},
			lengthDates: 0,
			err:         errors.New("error while parsing the ignored data, -r doesn't have arguments"),
		},
		{
			name:        "invalid option argument",
			args:        []string{"2020-04-03", "-a", "unexpected-arg"},
			lengthDates: 0,
			err:         errors.New("error while parsing the command found unknown argument"),
		},
		{
			name:        "more than one argument for format",
			args:        []string{"2020-04-03", "-f", "some format", "some second format"},
			lengthDates: 0,
			err:         errors.New("error while parsing the format, -f does have one argument"),
		},
		{
			name:        "some valid format",
			args:        []string{"2020-04-03", "2010-04-03", "-f", "some format"},
			lengthDates: 3654,
			err:         nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dates, err := GetDates(test.args)
			if test.err != nil {
				if err == nil {
					t.Fatalf("expected an error got nothing: expected: %v and got: %v \n", err, test.err)
				}
				if err.Error() != test.err.Error() {
					t.Fatalf("expected error: %v, got: %v", test.err, err)
				}
				return
			}
			if len(dates) <= 0 {
				t.Fatalf("expected not an empty date slice")
			}
			if test.lengthDates > 0 {
				if len(dates) != test.lengthDates {
					t.Fatalf("lenght dates did not match: expected: %v and got: %v \n", test.lengthDates, len(dates))
				}
			}
		})
	}
}

func TestGetAllDates(t *testing.T) {
	type datesTest struct {
		name        string
		args        []string
		lengthDates int
		err         error
	}
	var tests = []datesTest{
		{
			name:        "all normal",
			args:        []string{"2020-04-03"},
			lengthDates: -1,
			err:         nil,
		},
		{
			name:        "all normal, two arrays",
			args:        []string{"2020-04-03", "2020-05-03"},
			lengthDates: 31,
			err:         nil,
		},
		{
			name:        "no arguments provided",
			args:        []string{},
			lengthDates: 0,
			err:         errors.New("wrong number of dates provided"),
		},
		{
			name:        "more than two arguments provided",
			args:        []string{"2020-01-01", "2020-02-01", "2020-03-01"},
			lengthDates: 0,
			err:         errors.New("wrong number of dates provided"),
		},
		{
			name:        "invalid first date format",
			args:        []string{"invalid-date"},
			lengthDates: 0,
			err:         errors.New("first date provided is not valid"),
		},
		{
			name:        "invalid second date format",
			args:        []string{"2020-01-01", "invalid-date"},
			lengthDates: 0,
			err:         errors.New("second date provided is not valid"),
		},
		{
			name:        "first date equals second date",
			args:        []string{"2020-01-01", "2020-01-01"},
			lengthDates: 1, // Only one day expected
			err:         nil,
		},
		{
			name:        "first date after second date",
			args:        []string{"2020-02-01", "2020-01-01"},
			lengthDates: 0,   // This depends on your GetDatesFromTo behavior; if it returns empty slice or error
			err:         nil, // or an error if implemented
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dates, err := GetAllDates(test.args)
			if test.err != nil {
				if err == nil {
					t.Fatalf("expected an error got nothing: expected: %v and got: %v \n", err, test.err)
				}
				if err.Error() != test.err.Error() {
					t.Fatalf("expected error: %v, got: %v", test.err, err)
				}
				return
			}
			if len(dates) <= 0 {
				t.Fatalf("expected not an empty date slice")
			}
			if test.lengthDates > 0 {
				if len(dates) != test.lengthDates {
					t.Fatalf("lenght dates did not match: expected: %v and got: %v \n", test.lengthDates, len(dates))
				}
			}
		})
	}
}

func TestGetDatesFromTo(t *testing.T) {
	type datesFromToTest struct {
		name     string
		from     time.Time
		to       time.Time
		expected []time.Time
	}
	var datesFromToTests = []datesFromToTest{
		{
			"simple test",
			time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),
			time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC),
			[]time.Time{
				time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),
				time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC),
				time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			"simple test reversed",
			time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC),
			time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),
			[]time.Time{
				time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),
				time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC),
				time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			"same day",
			time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),
			time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),

			[]time.Time{
				time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, test := range datesFromToTests {
		t.Run(test.name, func(t *testing.T) {
			got := GetDatesFromTo(test.from, test.to)
			if len(test.expected) != len(got) {
				t.Fatalf("expected %d dates, got %d \n", len(test.expected), len(got))
			}
			for index, j := range test.expected {
				if j != got[index] {
					t.Fatalf("dates from to test mismatch at index %d, %v is not equal to %v \n", index, j.String(), got[index].String())
				}
			}
		})
	}
}

func TestRemoveWeekdays(t *testing.T) {
	// Create sample dates from Monday to Sunday
	dates := []time.Time{
		time.Date(2024, 6, 3, 0, 0, 0, 0, time.UTC), // Monday
		time.Date(2024, 6, 4, 0, 0, 0, 0, time.UTC), // Tuesday
		time.Date(2024, 6, 5, 0, 0, 0, 0, time.UTC), // Wednesday
		time.Date(2024, 6, 6, 0, 0, 0, 0, time.UTC), // Thursday
		time.Date(2024, 6, 7, 0, 0, 0, 0, time.UTC), // Friday
		time.Date(2024, 6, 8, 0, 0, 0, 0, time.UTC), // Saturday
		time.Date(2024, 6, 9, 0, 0, 0, 0, time.UTC), // Sunday
	}

	tests := []struct {
		name     string
		weekdays []time.Weekday
		allDates []time.Time
		want     []time.Time
	}{
		{
			name:     "Remove Monday and Tuesday",
			weekdays: []time.Weekday{time.Monday, time.Tuesday},
			allDates: dates,
			want: []time.Time{
				dates[2], // Wednesday
				dates[3], // Thursday
				dates[4], // Friday
				dates[5], // Saturday
				dates[6], // Sunday
			},
		},
		{
			name:     "Remove weekends",
			weekdays: []time.Weekday{time.Saturday, time.Sunday},
			allDates: dates,
			want: []time.Time{
				dates[0], // Monday
				dates[1], // Tuesday
				dates[2], // Wednesday
				dates[3], // Thursday
				dates[4], // Friday
			},
		},
		{
			name:     "Remove nothing",
			weekdays: []time.Weekday{},
			allDates: dates,
			want:     dates,
		},
		{
			name: "Remove all days",
			weekdays: []time.Weekday{
				time.Monday, time.Tuesday, time.Wednesday,
				time.Thursday, time.Friday, time.Saturday, time.Sunday,
			},
			allDates: dates,
			want:     []time.Time{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveWeekdays(tt.weekdays, tt.allDates)
			if len(got) != len(tt.want) {
				t.Fatalf("RemoveWeekdays() length = %d, want %d", len(got), len(tt.want))
			}
			for i := range got {
				if !got[i].Equal(tt.want[i]) {
					t.Errorf("RemoveWeekdays() got[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestIsNotInWeekdays(t *testing.T) {
	weekdays := []time.Weekday{time.Monday, time.Wednesday, time.Friday}

	tests := []struct {
		name     string
		day      time.Time
		expected bool
	}{
		{
			name:     "Day in weekdays (Monday)",
			day:      time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), // Monday
			expected: false,
		},
		{
			name:     "Day not in weekdays (Sunday)",
			day:      time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), // Sunday
			expected: true,
		},
		{
			name:     "Day in weekdays (Friday)",
			day:      time.Date(2023, 10, 6, 0, 0, 0, 0, time.UTC), // Friday
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNotInWeekdays(tt.day, weekdays)
			if result != tt.expected {
				t.Errorf("IsNotInWeekdays(%v) = %v, want %v", tt.day.Weekday(), result, tt.expected)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	input := []time.Time{
		time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 3, 0, 0, 0, 0, time.UTC),
	}
	expected := []time.Time{
		time.Date(2023, 10, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
	}

	result := Reverse(input)

	if len(result) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if !result[i].Equal(expected[i]) {
			t.Errorf("Reverse result[%d] = %v, want %v", i, result[i], expected[i])
		}
	}
}

func TestNeedHelp(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{
			name: "Contains help",
			args: []string{"foo", "bar", "-h"},
			want: true,
		},
		{
			name: "Help only",
			args: []string{"--help"},
			want: true,
		},
		{
			name: "No help",
			args: []string{"foo", "bar"},
			want: false,
		},
		{
			name: "Empty args",
			args: []string{},
			want: false,
		},
		{
			name: "Multiple helps",
			args: []string{"--help", "--help"},
			want: true,
		},
		{
			name: "Help as substring",
			args: []string{"helper", "helpful"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NeedHelp(tt.args)
			if got != tt.want {
				t.Errorf("NeedHelp(%v) = %v; want %v", tt.args, got, tt.want)
			}
		})
	}
}
func TestNeedVersion(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{
			name: "Contains version",
			args: []string{"foo", "bar", "-v"},
			want: true,
		},
		{
			name: "version only",
			args: []string{"--version"},
			want: true,
		},
		{
			name: "No version",
			args: []string{"foo", "bar"},
			want: false,
		},
		{
			name: "Empty args",
			args: []string{},
			want: false,
		},
		{
			name: "Multiple versions",
			args: []string{"--version", "--version"},
			want: true,
		},
		{
			name: "Help as substring",
			args: []string{"helper", "versionful"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NeedVersion(tt.args)
			if got != tt.want {
				t.Errorf("NeedHelp(%v) = %v; want %v", tt.args, got, tt.want)
			}
		})
	}
}

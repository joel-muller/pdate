package parser

import (
	"errors"
	"pdate/internal/constants"
	"pdate/internal/job"
	"reflect"
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantedJob job.Job
		wantErr   error
	}{
		{
			name: "only date",
			args: []string{"2025-03-4"},
			wantedJob: job.Job{
				DatesInput:      []time.Time{time.Date(2025, 03, 4, 0, 0, 0, 0, time.UTC)},
				PosArguments:    []job.Argument{job.Date},
				IgnoredWeekdays: []time.Weekday{},
				Format:          constants.DefaultInputFormat,
				Reversed:        false,
			},
			wantErr: nil,
		},
		{
			name: "nothing",
			args: []string{},
			wantedJob: job.Job{
				DatesInput:      []time.Time{},
				PosArguments:    []job.Argument{},
				IgnoredWeekdays: []time.Weekday{},
				Format:          constants.DefaultInputFormat,
				Reversed:        false,
			},
			wantErr: nil,
		},
		{
			name: "a lot of stuff",
			args: []string{"2025-03-18", "-f", "{MM}", "-i", "mo", "fr", "-r", "2022-08-10"},
			wantedJob: job.Job{
				DatesInput: []time.Time{
					time.Date(2025, 03, 18, 0, 0, 0, 0, time.UTC),
					time.Date(2022, 8, 10, 0, 0, 0, 0, time.UTC),
				},
				PosArguments:    []job.Argument{job.Date, job.Flag, job.Option, job.Flag, job.Option, job.Option, job.Flag, job.Date},
				IgnoredWeekdays: []time.Weekday{time.Monday, time.Friday},
				Format:          "{MM}",
				Reversed:        true,
			},
			wantErr: nil,
		},
		{
			name:      "ignore error",
			args:      []string{"2025-03-18", "-f", "{MM}", "-i", "mo", "wrong weekday", "-r", "2022-08-10"},
			wantedJob: job.Job{},
			wantErr:   errors.New("error while trying to parse a weekday"),
		},
		{
			name:      "error on sorting",
			args:      []string{"2025-03-18", "-a", "{MM}", "-i", "mo", "wrong weekday", "-r", "2022-08-10"},
			wantedJob: job.Job{},
			wantErr:   errors.New("found unknown flag"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := job.New()

			err := Parse(tt.args, j)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(j.DatesInput, tt.wantedJob.DatesInput) {
				t.Errorf("parsed job mismatch:\n got: %+v\nwant: %+v", j.DatesInput, tt.wantedJob.DatesInput)
			}
			if !reflect.DeepEqual(j.PosArguments, tt.wantedJob.PosArguments) {
				t.Errorf("parsed job mismatch:\n got: %+v\nwant: %+v", j.PosArguments, tt.wantedJob.PosArguments)
			}
			if !reflect.DeepEqual(j.IgnoredWeekdays, tt.wantedJob.IgnoredWeekdays) {
				t.Errorf("parsed job mismatch:\n got: %+v\nwant: %+v", j.IgnoredWeekdays, tt.wantedJob.IgnoredWeekdays)
			}
			if j.Format != tt.wantedJob.Format {
				t.Errorf("parsed job mismatch:\n got: %+v\nwant: %+v", j.Format, tt.wantedJob.Format)
			}
			if j.Reversed != tt.wantedJob.Reversed {
				t.Errorf("parsed job mismatch:\n got: %+v\nwant: %+v", j.Reversed, tt.wantedJob.Reversed)
			}
		})
	}
}

func TestParseInvalid(t *testing.T) {
	j := &job.Job{}

	err := ParseInvalid([]string{"anything"}, j)
	if err == nil {
		t.Errorf("expected error, got nothing")
	}
}

func TestParseHelp(t *testing.T) {
	j := &job.Job{}

	err := ParseHelp([]string{"anything"}, j)
	if err != nil {
		t.Errorf("error occured while trying to parse help")
	}
	if !j.Help {
		t.Errorf("help flag was not activated")
	}
}

func TestParseVersion(t *testing.T) {
	j := &job.Job{}

	err := ParseVersion([]string{"anything"}, j)
	if err != nil {
		t.Errorf("error occured while trying to parse version")
	}
	if !j.Version {
		t.Errorf("help flag was not activated")
	}
}

func TestParseFormat(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		initialJob job.Job
		wantFormat string
		wantErr    error
	}{
		{
			name:       "No arguments - returns error",
			args:       []string{},
			initialJob: job.Job{},
			wantFormat: "",
			wantErr:    errors.New("wrong format args given"),
		},
		{
			name:       "Multiple arguments - returns error",
			args:       []string{"json", "extra"},
			initialJob: job.Job{},
			wantFormat: "",
			wantErr:    errors.New("wrong format args given"),
		},
		{
			name:       "Single valid format",
			args:       []string{"json"},
			initialJob: job.Job{},
			wantFormat: "json",
			wantErr:    nil,
		},
		{
			name:       "Single valid format with different string",
			args:       []string{"custom-format"},
			initialJob: job.Job{},
			wantFormat: "custom-format",
			wantErr:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.initialJob
			err := ParseFormat(tt.args, &j)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if j.Format != tt.wantFormat {
				t.Errorf("expected Format to be %q, got %q", tt.wantFormat, j.Format)
			}
		})
	}
}

func TestParseIgnore(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		initialJob   job.Job
		wantWeekdays []time.Weekday
		wantErr      error
	}{
		{
			name:         "No arguments - returns error",
			args:         []string{},
			initialJob:   job.Job{},
			wantWeekdays: nil,
			wantErr:      errors.New("no weekdays for ignoring provided"),
		},
		{
			name:         "Invalid weekday - returns error",
			args:         []string{"xx"},
			initialJob:   job.Job{},
			wantWeekdays: nil,
			wantErr:      errors.New("error while trying to parse a weekday"),
		},
		{
			name:         "Single valid weekday",
			args:         []string{"mo"},
			initialJob:   job.Job{},
			wantWeekdays: []time.Weekday{time.Monday},
			wantErr:      nil,
		},
		{
			name:         "Multiple valid weekdays",
			args:         []string{"tu", "fr", "su"},
			initialJob:   job.Job{},
			wantWeekdays: []time.Weekday{time.Tuesday, time.Friday, time.Sunday},
			wantErr:      nil,
		},
		{
			name:         "Mixed with valid then invalid",
			args:         []string{"we", "bad"},
			initialJob:   job.Job{},
			wantWeekdays: []time.Weekday{time.Wednesday},
			wantErr:      errors.New("error while trying to parse a weekday"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.initialJob
			err := ParseIgnore(tt.args, &j)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(j.IgnoredWeekdays) != len(tt.wantWeekdays) {
				t.Errorf("expected weekdays %v, got %v", tt.wantWeekdays, j.IgnoredWeekdays)
				return
			}

			for i, wd := range tt.wantWeekdays {
				if j.IgnoredWeekdays[i] != wd {
					t.Errorf("expected weekday %v at index %d, got %v", wd, i, j.IgnoredWeekdays[i])
				}
			}
		})
	}
}

func TestParseReverse(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		initialJob   job.Job
		wantReversed bool
		wantErr      error
	}{
		{
			name:         "No arguments - sets reversed true",
			args:         []string{},
			initialJob:   job.Job{Reversed: false},
			wantReversed: true,
			wantErr:      nil,
		},
		{
			name:         "With arguments - returns error",
			args:         []string{"unexpected"},
			initialJob:   job.Job{Reversed: false},
			wantReversed: false,
			wantErr:      errors.New("reverse flag doesn't have arguments"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := tt.initialJob
			err := ParseReverse(tt.args, &j)

			if tt.wantErr != nil {
				if err == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if j.Reversed != tt.wantReversed {
				t.Errorf("expected Reversed to be %v, got %v", tt.wantReversed, j.Reversed)
			}
		})
	}
}

func TestSortOptions(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		expectErr    error
		expectSorted Sorted
	}{
		{
			name:      "Valid flags with values",
			args:      []string{"-i", "val1", "-f", "val2", "-r", "val3"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{
					Ignore:  {"val1"},
					Format:  {"val2"},
					Reverse: {"val3"},
				},
				dates:       []time.Time{},
				argumentPos: []job.Argument{job.Flag, job.Option, job.Flag, job.Option, job.Flag, job.Option},
			},
		},
		{
			name:      "Duplicate flag error",
			args:      []string{"-i", "val1", "-i", "val2"},
			expectErr: errors.New("found duplicate flag argument"),
		},
		{
			name:      "Unknown flag error",
			args:      []string{"-x", "oops"},
			expectErr: errors.New("found unknown flag"),
		},
		{
			name:      "Date parsing without flags",
			args:      []string{"2023-06-15"},
			expectErr: nil,
			expectSorted: Sorted{
				options:     map[flag][]string{},
				dates:       []time.Time{time.Date(2023, 6, 15, 0, 0, 0, 0, time.UTC)},
				argumentPos: []job.Argument{job.Date},
			},
		},
		{
			name:      "Mixed flag and date",
			args:      []string{"-i", "val", "2024-01-01"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{
					Ignore: {"val"},
				},
				dates:       []time.Time{time.Date(2024, 1, 01, 0, 0, 0, 0, time.UTC)},
				argumentPos: []job.Argument{job.Flag, job.Option, job.Date},
			},
		},
		{
			name:      "Mixed flag and two date",
			args:      []string{"-i", "2024-05-10", "val", "2024-01-01"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{
					Ignore: {"val"},
				},
				dates: []time.Time{
					time.Date(2024, 5, 10, 0, 0, 0, 0, time.UTC),
					time.Date(2024, 1, 01, 0, 0, 0, 0, time.UTC)},
				argumentPos: []job.Argument{job.Flag, job.Date, job.Option, job.Date},
			},
		},
		{
			name:      "No arguments",
			args:      []string{},
			expectErr: nil,
			expectSorted: Sorted{
				options:     map[flag][]string{},
				dates:       []time.Time{},
				argumentPos: []job.Argument{},
			},
		},
		{
			name:      "Only flags with no values",
			args:      []string{"-i", "-f", "-r"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{
					Ignore:  {},
					Format:  {},
					Reverse: {},
				},
				dates:       []time.Time{},
				argumentPos: []job.Argument{job.Flag, job.Flag, job.Flag},
			},
		},
		{
			name:      "Flag followed by multiple values",
			args:      []string{"-i", "val1", "val2", "-f", "val3"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{
					Ignore: {"val1", "val2"},
					Format: {"val3"},
				},
				dates:       []time.Time{},
				argumentPos: []job.Argument{job.Flag, job.Option, job.Option, job.Flag, job.Option},
			},
		},
		{
			name:      "Date in middle of values",
			args:      []string{"-i", "val1", "2025-01-01", "val2"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{
					Ignore: {"val1", "val2"},
				},
				dates:       []time.Time{time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)},
				argumentPos: []job.Argument{job.Flag, job.Option, job.Date, job.Option},
			},
		},
		{
			name:      "Date as value after unknown flag",
			args:      []string{"-x", "2025-01-01"},
			expectErr: errors.New("found unknown flag"),
		},
		{
			name:      "Flag at end of input",
			args:      []string{"val1", "-i"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{
					Ignore:  {},
					Invalid: {"val1"},
				},
				dates:       []time.Time{},
				argumentPos: []job.Argument{job.Option, job.Flag},
			},
		},
		{
			name:      "Multiple values, then date, then unknown flag",
			args:      []string{"-i", "v1", "2025-03-03", "-z"},
			expectErr: errors.New("found unknown flag"),
		},
		{
			name:      "Multiple dates only",
			args:      []string{"2024-01-01", "2024-12-31"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{},
				dates: []time.Time{
					time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
				},
				argumentPos: []job.Argument{job.Date, job.Date},
			},
		},
		{
			name:      "Flag with a date-like but invalid string",
			args:      []string{"-i", "2024-13-40", "val"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{
					Ignore: {"2024-13-40", "val"},
				},
				dates:       []time.Time{},
				argumentPos: []job.Argument{job.Flag, job.Option, job.Option},
			},
		},
		{
			name:      "Flag value is another flag-like string",
			args:      []string{"-i", "-notAFlag"},
			expectErr: errors.New("found unknown flag"),
		},
		{
			name:      "Version Flag",
			args:      []string{"-i", "2024-13-40", "val", "-v"},
			expectErr: nil,
			expectSorted: Sorted{
				options: map[flag][]string{
					Ignore:  {"2024-13-40", "val"},
					Version: {},
				},
				dates:       []time.Time{},
				argumentPos: []job.Argument{job.Flag, job.Option, job.Option, job.Flag},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SortOptions(tt.args)

			if tt.expectErr != nil {
				if err == nil || err.Error() != tt.expectErr.Error() {
					t.Errorf("expected error %v, got %v", tt.expectErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.expectSorted) {
				t.Errorf("expected sorted %v, got %v", tt.expectSorted, got)
			}
		})
	}
}

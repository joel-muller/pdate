package parser

import (
	"slices"
	"testing"
	"time"
)

func TestParseDate_ValidDate(t *testing.T) {
	dateStr := "2025-5-25"
	expected := time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC)

	result, err := ParseDate(dateStr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !result.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestParseDate_ValidDate_Padded(t *testing.T) {
	dateStr := "2025-05-25"
	expected := time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC)

	result, err := ParseDate(dateStr)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !result.Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
func TestParseDate_InvalidDateFormat(t *testing.T) {
	invalidDates := []string{
		"25-5-2025",
		"2025/05/25",
		"May 25, 2025",
		"",
		"hello world",
	}

	for _, dateStr := range invalidDates {
		_, err := ParseDate(dateStr)
		if err == nil {
			t.Errorf("Expected error for input %q, got nil", dateStr)
		}
	}
}

func TestParseDate_ZeroDate(t *testing.T) {
	dateStr := "0000-0-0"
	_, err := ParseDate(dateStr)
	if err == nil {
		t.Errorf("Expected error for zero date input, got nil")
	}
}

func TestParseIgnore(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []time.Weekday
	}{
		{
			name: "no -i flag",
			args: []string{"some", "args", "here"},
			want: []time.Weekday{},
		},
		{
			name: "-i with valid weekdays",
			args: []string{"cmd", "-i", "mo", "tu", "we"},
			want: []time.Weekday{time.Monday, time.Tuesday, time.Wednesday},
		},
		{
			name: "-i with invalid weekday stops parsing",
			args: []string{"-i", "mo", "xx", "we"},
			want: []time.Weekday{time.Monday},
		},
		{
			name: "-i at the end with no weekdays",
			args: []string{"foo", "bar", "-i"},
			want: []time.Weekday{},
		},
		{
			name: "-i with one valid weekday",
			args: []string{"-i", "fr"},
			want: []time.Weekday{time.Friday},
		},
		{
			name: "-i with valid weekdays followed by invalid",
			args: []string{"-i", "su", "sa", "xx"},
			want: []time.Weekday{time.Sunday, time.Saturday},
		},
		{
			name: "-i appears multiple times, parse only first",
			args: []string{"-i", "mo", "-i", "tu"},
			want: []time.Weekday{time.Monday},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseIgnore(tt.args)
			if !slices.Equal(tt.want, got) {
				t.Errorf("ParseIgnore() = %v, want %v", got, tt.want)
			}
		})
	}
}

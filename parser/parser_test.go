package parser

import (
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

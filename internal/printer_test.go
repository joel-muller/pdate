package internal

import (
	"testing"
	"time"
)

func TestPrintHelp(t *testing.T) {
	got := PrintHelp()
	want := helpMessage
	if got != want {
		t.Errorf("PrintHelp() = %q; want %q", got, want)
	}
}

func TestPrintVersion(t *testing.T) {
	got := PrintVersion()
	want := version
	if got != want {
		t.Errorf("PrintVersion() = %q; want %q", got, want)
	}
}

func TestDateFormatted(t *testing.T) {
	date := time.Date(2025, time.March, 5, 0, 0, 0, 0, time.UTC)

	t.Run("Default format", func(t *testing.T) {
		expected := "2025-03-05"
		result := DateFormatted(date, "")
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("Custom format", func(t *testing.T) {
		format := "{YYYY}/{MM}/{DD}"
		expected := "2025/03/05"
		result := DateFormatted(date, format)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}

func TestGetFormattedDates(t *testing.T) {
	dates := []time.Time{
		time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, time.December, 31, 0, 0, 0, 0, time.UTC),
	}
	format := "{YY}-{mn}-{D}"

	expected := []string{"25-Jan-1", "25-Dec-31"}
	result := GetFormattedDates(dates, format)

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("At index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
}

func TestReplaceDatePlaceholdersWithDate(t *testing.T) {
	date := time.Date(2025, time.March, 5, 0, 0, 0, 0, time.UTC)

	tests := map[string]string{
		"{YYYY}": "2025",
		"{YY}":   "25",
		"{MM}":   "03",
		"{M}":    "3",
		"{DD}":   "05",
		"{D}":    "5",
		"{WD}":   "Wednesday",
		"{wd}":   "Wed",
		"{MN}":   "March",
		"{mn}":   "Mar",
	}

	for placeholder, expected := range tests {
		result := ReplaceDatePlaceholdersWithDate(placeholder, date)
		if result != expected {
			t.Errorf("Placeholder %s: expected %s, got %s", placeholder, expected, result)
		}
	}

	t.Run("Full custom string", func(t *testing.T) {
		format := "Today is {WD}, {MN} {D}, {YYYY}"
		expected := "Today is Wednesday, March 5, 2025"
		result := ReplaceDatePlaceholdersWithDate(format, date)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}

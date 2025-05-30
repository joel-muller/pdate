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

func TestDateFormatted(t *testing.T) {
	// Define a fixed time to test formatting
	testDate := time.Date(2023, time.May, 30, 15, 0, 0, 0, time.UTC)

	got := DateFormatted(testDate, "")
	want := "30.05.2023"

	if got != want {
		t.Errorf("DateFormatted() = %q; want %q", got, want)
	}
}

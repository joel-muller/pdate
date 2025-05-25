package logic

import (
	"testing"
	"time"
)

func TestGetDatesFromTo_SimpleRange(t *testing.T) {
	from := time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC)

	expected := []time.Time{
		time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC),
	}

	got := GetDatesFromTo(from, to)

	if len(got) != len(expected) {
		t.Fatalf("expected %d dates, got %d", len(expected), len(got))
	}

	for i := range got {
		if !got[i].Equal(expected[i]) {
			t.Errorf("at index %d: expected %v, got %v", i, expected[i], got[i])
		}
	}
}

func TestGetDatesToFrom_SimpleRange(t *testing.T) {
	to := time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC)
	from := time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC)

	expected := []time.Time{
		time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 5, 24, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 5, 25, 0, 0, 0, 0, time.UTC),
	}

	got := GetDatesFromTo(from, to)

	if len(got) != len(expected) {
		t.Fatalf("expected %d dates, got %d", len(expected), len(got))
	}

	for i := range got {
		if !got[i].Equal(expected[i]) {
			t.Errorf("at index %d: expected %v, got %v", i, expected[i], got[i])
		}
	}
}

func TestGetDatesFromTo_SameDate(t *testing.T) {
	from := time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC)

	expected := []time.Time{
		time.Date(2025, 5, 23, 0, 0, 0, 0, time.UTC),
	}

	got := GetDatesFromTo(from, to)

	if len(got) != len(expected) {
		t.Fatalf("expected %d dates, got %d", len(expected), len(got))
	}

	for i := range got {
		if !got[i].Equal(expected[i]) {
			t.Errorf("at index %d: expected %v, got %v", i, expected[i], got[i])
		}
	}
}

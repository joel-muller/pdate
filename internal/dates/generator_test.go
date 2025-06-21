package dates

import (
	"testing"
	"time"
)

func TestGetDatesFromTo(t *testing.T) {
	tests := []struct {
		name         string
		from         time.Time
		to           time.Time
		expectedSize int
	}{
		{
			name:         "one day",
			from:         time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC),
			to:           time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC),
			expectedSize: 1,
		},
		{
			name:         "today",
			from:         time.Now(),
			to:           time.Now(),
			expectedSize: 1,
		},
		{
			name:         "one year",
			from:         time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC),
			to:           time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
			expectedSize: 366,
		},
		{
			name:         "one year reversed",
			from:         time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
			to:           time.Date(2025, 12, 1, 0, 0, 0, 0, time.UTC),
			expectedSize: 366,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetDatesFromTo(tt.from, tt.to)
			if len(result) != tt.expectedSize {
				t.Errorf("IsADayBefore() = %v, expected %v", len(result), tt.expectedSize)
			}
		})
	}
}

func TestIsADayBefore(t *testing.T) {
	tests := []struct {
		name     string
		before   time.Time
		after    time.Time
		expected bool
	}{
		{
			name:     "before is one day before after",
			before:   time.Date(2025, 6, 20, 15, 0, 0, 0, time.UTC),
			after:    time.Date(2025, 6, 21, 10, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "before and after are the same day",
			before:   time.Date(2025, 6, 21, 0, 0, 0, 0, time.UTC),
			after:    time.Date(2025, 6, 21, 23, 59, 59, 0, time.UTC),
			expected: false,
		},
		{
			name:     "before is after",
			before:   time.Date(2025, 6, 22, 0, 0, 0, 0, time.UTC),
			after:    time.Date(2025, 6, 21, 0, 0, 0, 0, time.UTC),
			expected: false,
		},
		{
			name:     "different months",
			before:   time.Date(2025, 5, 31, 23, 59, 59, 0, time.UTC),
			after:    time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "different years",
			before:   time.Date(2024, 12, 31, 12, 0, 0, 0, time.UTC),
			after:    time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: true,
		},
		{
			name:     "with timezones (before in UTC, after in PST)",
			before:   time.Date(2025, 6, 20, 23, 59, 59, 0, time.UTC),
			after:    time.Date(2025, 6, 21, 0, 0, 0, 0, time.FixedZone("PST", -8*3600)),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsADayBefore(tt.before, tt.after)
			if result != tt.expected {
				t.Errorf("IsADayBefore() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

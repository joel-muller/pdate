package dates

import (
	"testing"
	"time"
)

func TestRemoveWeekdays(t *testing.T) {
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
			got := IgnoreWeekdays(tt.allDates, tt.weekdays)
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

	result := ReverseOrder(input)

	if len(result) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(result))
	}

	for i := range expected {
		if !result[i].Equal(expected[i]) {
			t.Errorf("Reverse result[%d] = %v, want %v", i, result[i], expected[i])
		}
	}
}

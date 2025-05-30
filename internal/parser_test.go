package internal

import (
	"reflect"
	"testing"
	"time"
)

func TestParseIgnore(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []time.Weekday
		wantErr bool
	}{
		{
			name:    "valid weekdays",
			input:   []string{"mo", "tu", "we"},
			want:    []time.Weekday{time.Monday, time.Tuesday, time.Wednesday},
			wantErr: false,
		},
		{
			name:    "invalid weekday",
			input:   []string{"mo", "xy"},
			want:    []time.Weekday{},
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   []string{},
			want:    []time.Weekday{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseWeekdays(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseWeekdays() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseWeekdays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    time.Time
		wantErr bool
	}{
		{
			name:    "valid date",
			input:   "2023-12-25",
			want:    time.Date(2023, 12, 25, 0, 0, 0, 0, time.UTC),
			wantErr: false,
		},
		{
			name:    "invalid format",
			input:   "25-12-2023",
			want:    time.Time{},
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   "",
			want:    time.Time{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDate(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.want.IsZero() && got != tt.want {
				t.Errorf("ParseDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

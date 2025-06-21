package job

import (
	"pdate/internal/constants"
	"testing"
	"time"
)

// Helper to create a Job with custom arguments
func createJob(dates []time.Time, args []Argument, ignored []time.Weekday) *Job {
	return &Job{
		DatesInput:      dates,
		PosArguments:    args,
		IgnoredWeekdays: ignored,
		Reversed:        false,
		Format:          constants.DefaultInputFormat,
		Version:         false,
		Help:            false,
	}
}

func TestNew(t *testing.T) {
	j := New()
	if j == nil {
		t.Fatal("New() returned nil")
	}
	if len(j.DatesInput) != 0 {
		t.Error("Expected empty DatesInput")
	}
	if len(j.PosArguments) != 0 {
		t.Error("Expected empty PosArguments")
	}
	if len(j.IgnoredWeekdays) != 0 {
		t.Error("Expected empty IgnoredWeekdays")
	}
	if j.Format != constants.DefaultInputFormat {
		t.Error("Incorrect default format")
	}
	if j.Version {
		t.Error("Expected default Version to be false")
	}
	if j.Help {
		t.Error("Expected default Help to be false")
	}
}

func TestInvalidNumberOfDates(t *testing.T) {
	// Too many dates
	job := createJob([]time.Time{
		time.Now(), time.Now(), time.Now(),
	}, nil, nil)
	if !InvalidNumberOfDates(job) {
		t.Error("Expected true for too many dates")
	}

	// Valid number of dates
	job = createJob([]time.Time{
		time.Now(), time.Now(),
	}, nil, nil)
	if InvalidNumberOfDates(job) {
		t.Error("Expected false for 2 dates")
	}

	// No dates
	job = createJob([]time.Time{}, nil, nil)
	if InvalidNumberOfDates(job) {
		t.Error("Expected false for 0 dates")
	}
}

func TestDatesNextToEachOther(t *testing.T) {
	// Dates next to each other - OK
	job := createJob(nil, []Argument{Flag, Date, Date}, nil)
	if DatesNextToEachOther(job) {
		t.Error("Expected false for dates next to each other")
	}

	// Dates not next to each other
	job = createJob(nil, []Argument{Date, Flag, Date}, nil)
	if !DatesNextToEachOther(job) {
		t.Error("Expected true for dates not next to each other")
	}

	// Only one date
	job = createJob(nil, []Argument{Flag, Date, Flag}, nil)
	if DatesNextToEachOther(job) {
		t.Error("Expected false for single date")
	}
}

func TestDatesBetweenOptions(t *testing.T) {
	// Date followed by Option - invalid
	job := createJob(nil, []Argument{Flag, Date, Option}, nil)
	if !DatesBetweenOptions(job) {
		t.Error("Expected true for date followed by option")
	}

	// No date followed by option
	job = createJob(nil, []Argument{Flag, Date, Date}, nil)
	if DatesBetweenOptions(job) {
		t.Error("Expected false for dates not followed by option")
	}

	// No positional arguments
	job = createJob(nil, []Argument{}, nil)
	if DatesBetweenOptions(job) {
		t.Error("Expected false for no arguments")
	}
}

func TestDoubleWeekday(t *testing.T) {
	// Duplicate weekday
	job := createJob(nil, nil, []time.Weekday{time.Monday, time.Monday})
	if !DoubleWeekday(job) {
		t.Error("Expected true for duplicate weekdays")
	}

	// Unique weekdays
	job = createJob(nil, nil, []time.Weekday{time.Monday, time.Tuesday})
	if DoubleWeekday(job) {
		t.Error("Expected false for unique weekdays")
	}

	// Empty weekdays
	job = createJob(nil, nil, []time.Weekday{})
	if DoubleWeekday(job) {
		t.Error("Expected false for empty weekdays")
	}
}

func TestValidate(t *testing.T) {
	// Too many dates
	job := createJob([]time.Time{time.Now(), time.Now(), time.Now()}, nil, nil)
	err := Validate(job)
	if err == nil || err.Error() != "wrong number of dates provided" {
		t.Error("Expected wrong number of dates error")
	}

	// Dates not next to each other
	job = createJob([]time.Time{time.Now(), time.Now()}, []Argument{Date, Flag, Date}, nil)
	err = Validate(job)
	if err == nil || err.Error() != "the two dates are not next to each other" {
		t.Error("Expected 'dates not next to each other' error")
	}

	// Dates between options
	job = createJob([]time.Time{time.Now(), time.Now()}, []Argument{Date, Option}, nil)
	err = Validate(job)
	if err == nil || err.Error() != "dates are in between options" {
		t.Error("Expected 'dates between options' error")
	}

	// Double weekday
	job = createJob([]time.Time{time.Now(), time.Now()}, []Argument{Date, Date}, []time.Weekday{time.Monday, time.Monday})
	err = Validate(job)
	if err == nil || err.Error() != "double weekdays for ignore -i flag detected" {
		t.Error("Expected 'double weekday' error")
	}

	// All valid
	job = createJob([]time.Time{time.Now(), time.Now()}, []Argument{Date, Date}, []time.Weekday{time.Monday})
	err = Validate(job)
	if err != nil {
		t.Errorf("Expected nil error for valid job, got: %v", err)
	}
}

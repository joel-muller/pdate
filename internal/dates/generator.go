package dates

import (
	"pdate/internal/constants"
	"pdate/internal/job"
	"time"
)

func GetDates(j *job.Job) []string {
	if j.Help {
		return []string{constants.HelpMessage}
	}
	if j.Version {
		return []string{constants.Version}
	}
	allDates := GetAllDates(j.DatesInput)
	ignoredWeekdays := IgnoreWeekdays(allDates, j.IgnoredWeekdays)
	if j.Reversed {
		ignoredWeekdays = ReverseOrder(ignoredWeekdays)
	}
	return FormatDates(ignoredWeekdays, j.Format, j.Language)
}

func GetAllDates(dates []time.Time) []time.Time {
	switch len(dates) {
	case 2:
		return GetDatesFromTo(dates[0], dates[1])
	case 1:
		return GetDatesFromTo(dates[0], time.Now())
	default:
		return GetDatesFromTo(time.Now(), time.Now())
	}
}

func GetDatesFromTo(from time.Time, to time.Time) []time.Time {
	var lower = from
	var upper = to
	if upper.Before(lower) {
		upper, lower = lower, upper
	}
	var dates []time.Time
	dates = append(dates, lower)
	for IsADayBefore(lower, upper) {
		lower = lower.Add(time.Hour * 24)
		dates = append(dates, lower)
	}
	return dates
}

func IsADayBefore(before time.Time, after time.Time) bool {
	beforeY, beforeM, beforeD := before.Date()
	afterY, afterM, afterD := after.Date()
	beforeDate := time.Date(beforeY, beforeM, beforeD, 0, 0, 0, 0, time.UTC)
	afterDate := time.Date(afterY, afterM, afterD, 0, 0, 0, 0, time.UTC)
	return beforeDate.Before(afterDate)
}

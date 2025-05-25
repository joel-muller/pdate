package logic

import (
	"errors"
	"pdate/parser"
	"time"
)

func GetDates(args []string) ([]time.Time, error) {
	dates, err := GetAllDates(args)
	if err != nil {
		return []time.Time{}, err
	}
	dates = RemoveWeekdays(args, dates)
	return dates, nil
}

func GetAllDates(args []string) ([]time.Time, error) {
	dates := parser.GetDates(args)
	if len(dates) == 2 {
		return GetDatesFromTo(dates[0], dates[1]), nil
	}
	if len(dates) == 1 {
		return GetDatesFromTo(dates[0], time.Now()), nil
	}
	return []time.Time{}, errors.New("didnt got 2 or 1 valid date")
}

func GetDatesFromTo(from time.Time, to time.Time) []time.Time {
	var dates []time.Time
	var lower = from
	var upper = to
	if upper.Before(lower) {
		upper = from
		lower = to
	}
	dates = append(dates, lower)
	for lower.Before(upper) {
		lower = lower.Add(time.Hour * 24)
		dates = append(dates, lower)
	}
	return dates
}

func RemoveWeekdays(args []string, allDates []time.Time) []time.Time {
	weekdays := parser.ParseIgnore(args)
	var dates []time.Time
	for _, date := range allDates {
		if IsNotInWeekdays(date, weekdays) {
			dates = append(dates, date)
		}
	}
	return dates
}

func IsNotInWeekdays(day time.Time, weekdays []time.Weekday) bool {
	for _, weekday := range weekdays {
		if day.Weekday() == weekday {
			return false
		}
	}
	return true
}

package internal

import (
	"errors"
	"fmt"
	"time"
)

func NeedHelp(args []string) bool {
	for _, i := range args {
		if i == "-h" || i == "--help" {
			return true
		}
	}
	return false
}

func GetDates(args []string) ([]time.Time, error) {
	sorted, sortedErr := SortArguments(args)
	if sortedErr != nil {
		return []time.Time{}, fmt.Errorf("error while parsing the command %v", sortedErr)
	}
	dates, datesError := GetAllDates(sorted[DATE])
	if datesError != nil {
		return []time.Time{}, fmt.Errorf("error while parsing dates %v", datesError)
	}
	ignore, hasIgnore := sorted[IGNORE]
	if hasIgnore {
		weekdays, ignoreError := ParseWeekdays(ignore)
		if ignoreError != nil {
			return []time.Time{}, fmt.Errorf("error while parsing the ignored data %v", ignoreError)
		}
		ignoredDates := RemoveWeekdays(weekdays, dates)
		dates = ignoredDates
	}
	reverse, hasReverse := sorted[REVERSE]
	if hasReverse {
		if len(reverse) != 0 {
			return []time.Time{}, fmt.Errorf("error while parsing the ignored data, -r doesn't have arguments")
		}
		dates = Reverse(dates)
	}
	return dates, nil
}

func GetAllDates(args []string) ([]time.Time, error) {
	if len(args) < 1 || len(args) > 2 {
		return []time.Time{}, errors.New("wrong number of dates provided")
	}
	date1, err1 := ParseDate(args[0])
	if err1 != nil {
		return []time.Time{}, errors.New("first date provided is not valid")
	}
	if len(args) == 2 {
		date2, err2 := ParseDate(args[1])
		if err2 != nil {
			return []time.Time{}, errors.New("second date provided is not valid")
		}
		return GetDatesFromTo(date1, date2), nil
	}
	return GetDatesFromTo(date1, time.Now()), nil
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

func RemoveWeekdays(weekdays []time.Weekday, allDates []time.Time) []time.Time {
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

func Reverse(array []time.Time) []time.Time {
	var reversed []time.Time
	for i := len(array) - 1; i >= 0; i-- {
		reversed = append(reversed, array[i])
	}
	return reversed
}

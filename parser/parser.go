package parser

import (
	"errors"
	"time"
)

const (
	IGNORE = "-i"
)

var strToWeekday = map[string]time.Weekday{
	"mo": time.Monday,
	"tu": time.Tuesday,
	"we": time.Wednesday,
	"th": time.Thursday,
	"fr": time.Friday,
	"sa": time.Saturday,
	"su": time.Sunday,
}

func ParseDate(date string) (time.Time, error) {
	layout := "2006-1-2"
	t, err := time.Parse(layout, date)
	if err != nil || t.IsZero() {
		return time.Time{}, errors.New("could not parse the date")
	}
	return t, nil
}

func GetDates(args []string) []time.Time {
	var dates []time.Time
	for _, arg := range args {
		date, err := ParseDate(arg)
		if err == nil {
			dates = append(dates, date)
		}
	}
	return dates
}

func ParseIgnore(args []string) []time.Weekday {
	var days []time.Weekday
	var index = -1
	for i, arg := range args {
		if arg == IGNORE {
			index = i
			break
		}
	}
	if index == -1 {
		return days
	}
	for i := index + 1; i < len(args); i++ {
		day, exists := strToWeekday[args[i]]
		if exists {
			days = append(days, day)
		} else {
			return days
		}
	}
	return days
}

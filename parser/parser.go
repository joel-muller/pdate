package parser

import (
	"errors"
	"time"
)

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

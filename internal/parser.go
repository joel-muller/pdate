package internal

import (
	"errors"
	"time"
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

func ParseWeekdays(args []string) ([]time.Weekday, error) {
	var days []time.Weekday
	if len(args) == 0 {
		return []time.Weekday{}, errors.New("error, got ignore filed -i but got nothing to ignore")
	}
	for _, dayString := range args {
		day, exists := strToWeekday[dayString]
		if exists {
			days = append(days, day)
		} else {
			return []time.Weekday{}, errors.New("error, could not parse the weekdays arguments correctly")
		}
	}
	return days, nil
}

func ParseDate(date string) (time.Time, error) {
	layout := "2006-1-2"
	t, err := time.Parse(layout, date)
	if err != nil || t.IsZero() {
		return time.Time{}, errors.New("could not parse the date")
	}
	return t, nil
}

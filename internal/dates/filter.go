package dates

import (
	"time"
)

func IgnoreWeekdays(dates []time.Time, weekdays []time.Weekday) []time.Time {
	weekdayMap := make(map[time.Weekday]bool)
	for _, w := range weekdays {
		weekdayMap[w] = true
	}
	var result []time.Time
	for _, date := range dates {
		if !weekdayMap[date.Weekday()] {
			result = append(result, date)
		}
	}
	return result
}

func ReverseOrder(array []time.Time) []time.Time {
	var reversed []time.Time
	for i := len(array) - 1; i >= 0; i-- {
		reversed = append(reversed, array[i])
	}
	return reversed
}

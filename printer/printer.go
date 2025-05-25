package printer

import (
	"fmt"
	"time"
)

const (
	YEAR  = "%Y"
	MONTH = "%m"
	DAY   = "%d"
)

func PrintDates(dates []time.Time, format string) {
	for _, date := range dates {
		fmt.Println(DateFormatted(date, format))
	}
}

// DateFormatted Default, will change later
func DateFormatted(date time.Time, format string) string {
	return date.Format("02.01.2006")
}

package dates

import (
	"fmt"
	"strings"
	"time"
)

func FormatDates(dates []time.Time, format string) []string {
	var formattedDates []string
	for _, date := range dates {
		formattedDates = append(formattedDates, ReplaceDatePlaceholdersWithDate(format, date))
	}
	return formattedDates
}

func ReplaceDatePlaceholdersWithDate(input string, date time.Time) string {
	replacer := strings.NewReplacer(
		"{YYYY}", fmt.Sprintf("%04d", date.Year()),
		"{YY}", fmt.Sprintf("%02d", date.Year()%100),
		"{MM}", fmt.Sprintf("%02d", int(date.Month())),
		"{DD}", fmt.Sprintf("%02d", date.Day()),
		"{WD}", date.Weekday().String(),
		"{wd}", date.Weekday().String()[:3],
		"{MN}", date.Month().String(),
		"{mn}", date.Month().String()[:3],
		"{M}", fmt.Sprintf("%d", int(date.Month())),
		"{D}", fmt.Sprintf("%d", date.Day()),
	)
	return replacer.Replace(input)
}

package internal

import (
	"fmt"
	"strings"
	"time"
)

const (
	YEAR  = "%Y"
	MONTH = "%m"
	DAY   = "%d"
)

const helpMessage = `Usage:
  pdate <start_date> [end_date] [-i <days_to_exclude>] [-r]

Description:
  Prints dates from the start_date up to end_date (or today if end_date is omitted).

Options:
  <start_date>         Start date in YYYY-MM-DD format.
  [end_date]           Optional end date in YYYY-MM-DD format.
  -i <days_to_exclude> Exclude specified weekdays (e.g., mo, tu, fr).
  -r                   Print dates in reverse order.

Examples:
  pdate 2025-10-02
    Prints all dates from October 2, 2025 to today.

  pdate 2025-10-02 2025-11-30
    Prints all dates from October 2 to November 30, 2025.

  pdate 2025-10-02 2025-11-30 -i mo tu
    Prints dates excluding Mondays and Tuesdays.

  pdate 2025-10-02 2025-11-30 -i mo tu fr sa su -r
    Prints dates excluding Mon, Tue, Fri, Sat, Sun in reverse order.
`

func PrintHelp() string {
	return helpMessage
}

func DateFormatted(date time.Time, format string) string {
	if len(format) > 0 {
		return ReplaceDatePlaceholdersWithDate(format, date)
	}
	return date.Format("2006-01-02")
}

func GetFormattedDates(dates []time.Time, format string) []string {
	var formattedDates []string
	for _, date := range dates {
		formattedDates = append(formattedDates, DateFormatted(date, format))
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

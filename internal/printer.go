package internal

import (
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

// DateFormatted Default, will change later
func DateFormatted(date time.Time, format string) string {
	return date.Format("02.01.2006")
}

package internal

import (
	"fmt"
	"strings"
	"time"
)

const version = "1.2.0"

const helpMessage = `Usage:
  pdate <start-date> [end-date] [-i <days-to-ignore>] [-f <format>] [-r]

Description:
  Prints dates from <start-date> to <end-date> (or today if end-date is omitted).
  You can optionally ignore specific weekdays, customize the date format, or reverse the order.

Options:
  <start-date>         Start of the date range (format: YYYY-MM-DD).
  [end-date]           Optional end of the range (format: YYYY-MM-DD). Defaults to today.
  -i <days>            Ignore specific weekdays using codes (e.g., mo tu fr).
  -f <format>          Format each date using placeholders (see below).
  -r                   Print dates in reverse order.
  -h, --help           Show this help message.
  -v, --version        Show version

Weekday Codes for -i:
  mo  Monday
  tu  Tuesday
  we  Wednesday
  th  Thursday
  fr  Friday
  sa  Saturday
  su  Sunday

Format Placeholders for -f:
  {YYYY}  Full year (e.g., 2025)
  {YY}    Last two digits of year (e.g., 25)
  {MM}    Month with leading zero (e.g., 12)
  {M}     Month without leading zero (e.g., 12)
  {DD}    Day with leading zero (e.g., 07)
  {D}     Day without leading zero (e.g., 7)
  {MN}    Full month name (e.g., December)
  {mn}    Abbreviated month name (e.g., Dec)
  {WD}    Full weekday name (e.g., Sunday)
  {wd}    Abbreviated weekday name (e.g., Sun)

Examples:
  pdate 2025-10-02
    Prints all dates from October 2, 2025 to today.

  pdate 2025-10-02 2025-11-30
    Prints all dates from October 2 to November 30, 2025.

  pdate 2025-10-02 2025-11-30 -i mo tu
    Prints dates excluding Mondays and Tuesdays.

  pdate 2025-10-02 2025-11-30 -i mo tu fr sa su -r
    Prints dates excluding Mon, Tue, Fri, Sat, Sun in reverse order.

  pdate 2025-10-02 2025-10-10 -f "{DD}.{MM}.{YYYY} ({wd})"
    Prints formatted dates like 02.10.2025 (Thu)
`

func PrintHelp() string {
	return helpMessage
}

func PrintVersion() string {
	return version
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

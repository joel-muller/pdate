package constants

const Version = "1.2.0"

const DefaultInputFormat = "{YYYY}-{MM}-{DD}"

const ParseLayoutDate = "2006-1-2"

const HelpMessage = `Usage:
  pdate [-i <days-to-ignore>] [-f <format>] [-r] [-l <language>] [start-date] [end-date]

Description:
  Prints dates from <start-date> to <end-date> (or today if end-date is omitted).
  You can optionally ignore specific weekdays, customize the date format, or reverse the order.

Options:
  [start-date]         Start of the date range (format: YYYY-MM-DD).
  [end-date]           Optional end of the range (format: YYYY-MM-DD). Defaults to today.
  -i <days>            Ignore specific weekdays using codes (e.g., mo tu fr).
  -f <format>          Format each date using placeholders (see below).
  -r                   Print dates in reverse order.
  -l <language>        Print the format in the desired language, default language is english
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

Language Codes for -l:
  en  English
  fr  French
  es  Spanish
  de  German
  ch  Swiss
  it  Italian
  pt  Portuguese
  nl  Dutch
  ru  Russian
  zh  Chinese
  ar  Arabic
  hi  Hindi

Examples:
  pdate 2025-10-02
    Prints all dates from October 2, 2025 to today.

  pdate 2025-10-02 2025-11-30
    Prints all dates from October 2 to November 30, 2025.

  pdate -i mo tu 2025-10-02 2025-11-30
    Prints dates excluding Mondays and Tuesdays.

  pdate -i mo tu fr sa su -r 2025-10-02 2025-11-30
    Prints dates excluding Mon, Tue, Fri, Sat, Sun in reverse order.

  pdate -f "{DD}.{MM}.{YYYY} ({wd})" 2025-10-02 2025-10-10
    Prints formatted dates like 02.10.2025 (Thu)
`

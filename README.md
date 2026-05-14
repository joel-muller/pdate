# pdate - A simple date printer 📆

`pdate` is a command-line utility that prints a range of dates between two specified days. It supports options to ignore specific weekdays and to reverse the output order.

[![Go](https://github.com/joel-muller/pdate/actions/workflows/build.yml/badge.svg)](https://github.com/joel-muller/pdate/actions/workflows/build.yml) [![goreleaser](https://github.com/joel-muller/pdate/actions/workflows/release.yml/badge.svg)](https://github.com/joel-muller/pdate/actions/workflows/release.yml)

## Installation

### Homebrew

If you use Homebrew, you can install the app via the official cask:

```bash
brew install --cask joel-muller/tap/pdate
```

### Other platforms

For all other devices, download the latest version from the [pdate releases page](https://github.com/joel-muller/pdate/releases)

## Usage

```bash
pdate [-i <days-to-ignore>] [-f <format>] [-r] [-l <language>] [start-date] [end-date]
```

* `start-date`: The beginning of the date range (format: `YYYY-MM-DD`)
* `end-date`: *(Optional)* The end of the date range (format: `YYYY-MM-DD`). If omitted, the range ends at **today's date**.
* `-i <days>`: *(Optional)* Ignore specific weekdays. You can list one or more weekday codes after `-i`. 
* `-f <format>`: *(Optional)* Format the date in a provided format (listed after `-i` between two `""`) in a string (see bellow)
* `-r`: *(Optional)* Print the resulting list of dates in reverse order.
* `-l <language>`: *(Optional)* Print the format in the desired language, default language is english
* `-h` or `--help`: Display help information about `pdate`
* `-v` or `--version`: Display the version of `pdate`

### Weekday Codes

Use these short codes with the `-i` flag to ignore specific weekdays:

| Code | Day       |
|------|-----------|
| `mo` | Monday    |
| `tu` | Tuesday   |
| `we` | Wednesday |
| `th` | Thursday  |
| `fr` | Friday    |
| `sa` | Saturday  |
| `su` | Sunday    |

### Format Placeholders

Use the `-f` flag with these placeholders to customize date output:

| Placeholder | Description                       | Example (`2025-12-07`) |
|-------------|-----------------------------------|------------------------|
| `{YYYY}`    | Full year                         | `2025`                 |
| `{YY}`      | Last two digits of the year       | `25`                   |
| `{MM}`      | Month with leading zero           | `12`                   |
| `{M}`       | Month without leading zero        | `12`                   |
| `{DD}`      | Day of month with leading zero    | `07`                   |
| `{D}`       | Day of month without leading zero | `7`                    |
| `{MN}`      | Full month name                   | `December`             |
| `{mn}`      | Abbreviated month name            | `Dec`                  |
| `{WD}`      | Full weekday name                 | `Sunday`               |
| `{wd}`      | Abbreviated weekday name          | `Sun`                  |

### Language Codes

Use these short country codes for the `-l` flag and `{MN}`, `{mn}`, `{WD}` and `{wd}` is parsed in the desired language

| Code | Language        |
|------|-----------------|
| `en` | 🇬🇧 English    |
| `fr` | 🇫🇷 French     |
| `es` | 🇪🇸 Spanish    |
| `de` | 🇩🇪 German     |
| `ch` | 🇨🇭 Swiss      |
| `it` | 🇮🇹 Italian    |
| `pt` | 🇵🇹 Portuguese |
| `nl` | 🇳🇱 Dutch      |
| `ru` | 🇷🇺 Russian    |
| `zh` | 🇨🇳 Chinese    |
| `ar` | 🇸🇦 Arabic     |
| `hi` | 🇮🇳 Hindi      |

### Example Commands

```bash
pdate 2025-10-02
```

> Prints all dates from October 2, 2025 to today.

```bash
pdate 2025-10-02 2025-11-30
```

> Prints all dates from October 2 to November 30, 2025.

```bash
pdate -i mo tu fr sa su -r 2025-10-02 2025-11-30
```

> Prints dates from the same range, **excluding Mon, Tue, Fri, Sat, Sun**, and prints them in **reverse order**.

```bash
pdate -f "{YYYY}-{MM}-{DD} ({WD})" 2025-10-02 2025-11-05
```

> Prints dates in a custom format, e.g., `2025-10-02 (Thursday)`.

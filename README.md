# pdate - A simple date printer 📆

`pdate` is a command-line utility that prints a range of dates between two specified days. It supports options to ignore specific weekdays and to reverse the output order.

[![Go](https://github.com/joel-muller/pdate/actions/workflows/go.yml/badge.svg)](https://github.com/joel-muller/pdate/actions/workflows/go.yml) [![goreleaser](https://github.com/joel-muller/pdate/actions/workflows/release.yml/badge.svg)](https://github.com/joel-muller/pdate/actions/workflows/release.yml)

## Usage

```bash
pdate <start-date> [end-date] [-i <days-to-ignore>] [-f <format>] [-r]
```

* `start-date`: The beginning of the date range (format: `YYYY-MM-DD`)
* `end-date`: *(Optional)* The end of the date range (format: `YYYY-MM-DD`). If omitted, the range ends at **today's date**.
* `-i <days>`: *(Optional)* Ignore specific weekdays. You can list one or more weekday codes after `-i`. 
* `-f <format>`: *(Optional)* Format the date in a provided format (listed after `-i` between two `""`) in a string (see bellow)
* `-r`: *(Optional)* Print the resulting list of dates in reverse order.
* `-h` or `--help`: Display help information about `pdate`


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
pdate 2025-10-02 2025-11-30 -i mo tu
```

> Prints all dates from October 2 to November 30, 2025, **excluding Mondays and Tuesdays**.

```bash
pdate 2025-10-02 2025-11-30 -i mo tu fr sa su -r
```

> Prints dates from the same range, **excluding Mon, Tue, Fri, Sat, Sun**, and prints them in **reverse order**.

```bash
pdate 2025-10-02 2025-11-05 -f "{YYYY}-{MM}-{DD} ({WD})"
```

> Prints dates in a custom format, e.g., `2025-10-02 (Thursday)`.

## Installation

### Linux

1. Download the archive (e.g. `pdate_1.1.0_linux_amd64.tar.gz`)

   <details><summary>Optional: Check checksum</summary>

   ```bash
   sha256sum -c pdate_1.1.0_checksums.txt
   ```

   </details>

2. Extract it:

   ```bash
   tar -xzf pdate_1.1.0_linux_amd64.tar.gz
   ```

3. Move it to your system path and make it executable:

   ```bash
   sudo mv pdate /usr/local/bin/
   ```

4. Run it:

   ```bash
   pdate 2025-10-02
   ```

### macOS

1. Download the archive (e.g. `pdate_1.1.0_darwin_arm64.tar.gz`)

   <details><summary>Optional: Check checksum</summary>

   ```bash
   shasum -a 256 -c pdate_1.1.0_checksums.txt
   ```

   </details>

2. Extract it:

   ```bash
   tar -xzf pdate_1.1.0_darwin_arm64.tar.gz
   ```

3. Move it to your system path:

   ```bash
   sudo mv pdate /usr/local/bin/
   ```

4. Run it:

   ```bash
   pdate 2025-10-02
   ```

> [!WARNING]
> On macOS, you might need to allow the app to run the first time:
> Go to System Settings → Privacy & Security → Security and click “Allow Anyway” if macOS blocks the binary.

### Windows

I don’t personally use Windows, but a Windows binary is available. If you know how to install and run `pdate` on Windows, please feel free to update this section and submit a pull request. Contributions are always welcome!
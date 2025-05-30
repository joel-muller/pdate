# pdate - A simple date printer 📆

`pdate` is a command-line utility that prints a range of dates between two specified days. It supports options to ignore specific weekdays and to reverse the output order.

[![Go](https://github.com/joel-muller/pdate/actions/workflows/go.yml/badge.svg)](https://github.com/joel-muller/pdate/actions/workflows/go.yml) [![goreleaser](https://github.com/joel-muller/pdate/actions/workflows/release.yml/badge.svg)](https://github.com/joel-muller/pdate/actions/workflows/release.yml)

## Usage

```bash
pdate <start-date> [end-date] [-i <days-to-ignore>] [-r]
```

* `start-date`: The beginning of the date range (format: `YYYY-MM-DD`)
* `end-date`: *(Optional)* The end of the date range (format: `YYYY-MM-DD`). If omitted, the range ends at **today's date**.
* `-i <days>`: *(Optional)* Ignore specific weekdays. You can list one or more weekday codes after `-i`.
* `-r`: *(Optional)* Print the resulting list of dates in reverse order.
* `-h` or `--help`: Display help information about `pdate`

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

### Weekday Codes

Use these short codes with the `-i` flag to ignore specific weekdays:

| Code | Day       |
|------|-----------|
| mo   | Monday    |
| tu   | Tuesday   |
| we   | Wednesday |
| th   | Thursday  |
| fr   | Friday    |
| sa   | Saturday  |
| su   | Sunday    |

## Installation

### Linux

1. Download the archive (e.g. `pdate_1.0.0_linux_amd64.tar.gz`) and check if the checksum matches the binary

   ```bash
   sha256sum -c pdate_1.0.0_checksums.txt
   ```

2. Extract it:

   ```bash
   tar -xzf pdate_1.0.0_linux_amd64.tar.gz
   ```
3. Move it to your system path and make it executable:

   ```bash
   sudo mv pdate /usr/local/bin/
   chmod +x /usr/local/bin/pdate
   ```
4. Run it:

   ```bash
   pdate 2025-10-02
   ```

### macOS

1. Download the archive (e.g. `pdate_1.0.0_darwin_arm64.tar.gz`) and check if the checksum matches the binary

   ```bash
   shasum -a 256 -c pdate_1.0.0_checksums.txt
   ```

2. Extract it:

   ```bash
   tar -xzf pdate_1.0.0_darwin_arm64.tar.gz
   ```
3. Move it to your system path:

   ```bash
   sudo mv pdate /usr/local/bin/
   chmod +x /usr/local/bin/pdate
   ```
4. Run it:

   ```bash
   pdate 2025-10-02
   ```

> Note: On macOS, you might need to allow the app to run the first time:
Go to System Settings → Privacy & Security → Security and click “Allow Anyway” if macOS blocks the binary.

### Windows

I don’t personally use Windows, but a Windows binary is available. If you know how to install and run `pdate` on Windows, please feel free to update this section and submit a pull request. Contributions are always welcome!


## Future Updates

* Format output for better readability
* Support multiple input date formats (e.g., `YYYY-MM-DD`, `MM/DD/YYYY`, etc.)
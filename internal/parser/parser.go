package parser

import (
	"errors"
	"pdate/internal/constants"
	"pdate/internal/job"
	"time"
)

type Sorted struct {
	options     map[flag][]string
	dates       []time.Time
	argumentPos []job.Argument
}

type flag int

const (
	Ignore flag = iota
	Reverse
	Format
	Version
	Help
	Invalid
)

var strToOption = map[string]flag{
	"-i":        Ignore,
	"-r":        Reverse,
	"-f":        Format,
	"-v":        Version,
	"--version": Version,
	"-h":        Help,
	"--help":    Help,
}

var optionToJobFunc = map[flag]func([]string, *job.Job) error{
	Ignore:  ParseIgnore,
	Reverse: ParseReverse,
	Format:  ParseFormat,
	Version: ParseVersion,
	Help:    ParseHelp,
	Invalid: ParseInvalid,
}

var strToWeekday = map[string]time.Weekday{
	"mo": time.Monday,
	"tu": time.Tuesday,
	"we": time.Wednesday,
	"th": time.Thursday,
	"fr": time.Friday,
	"sa": time.Saturday,
	"su": time.Sunday,
}

func Parse(args []string, job *job.Job) error {
	sorted, err := SortOptions(args)
	if err != nil {
		return err
	}
	job.DatesInput = sorted.dates
	job.PosArguments = sorted.argumentPos
	for key, value := range sorted.options {
		parseMethod, found := optionToJobFunc[key]
		if found {
			jobError := parseMethod(value, job)
			if jobError != nil {
				return jobError
			}
		}
	}
	return nil
}

func ParseInvalid(args []string, job *job.Job) error {
	return errors.New("invalid option given for no flag")
}

func ParseIgnore(args []string, job *job.Job) error {
	if len(args) == 0 {
		return errors.New("no weekdays for ignoring provided")
	}
	for _, arg := range args {
		weekday, valid := strToWeekday[arg]
		if !valid {
			return errors.New("error while trying to parse a weekday")
		}
		job.IgnoredWeekdays = append(job.IgnoredWeekdays, weekday)
	}
	return nil
}

func ParseHelp(args []string, job *job.Job) error {
	job.Help = true
	return nil
}

func ParseVersion(args []string, job *job.Job) error {
	job.Version = true
	return nil
}

func ParseFormat(args []string, job *job.Job) error {
	if len(args) != 1 {
		return errors.New("wrong format args given")
	}
	job.Format = args[0]
	return nil
}

func ParseReverse(args []string, job *job.Job) error {
	if len(args) != 0 {
		return errors.New("reverse flag doesn't have arguments")
	}
	job.Reversed = true
	return nil
}

func SortOptions(args []string) (Sorted, error) {
	sorted := Sorted{
		map[flag][]string{},
		[]time.Time{},
		[]job.Argument{},
	}
	var currentOption = Invalid
	for _, arg := range args {
		if len(arg) > 0 && arg[0] == '-' {
			newOption, val := strToOption[arg]
			if !val {
				return Sorted{}, errors.New("found unknown flag")
			}
			_, isAlreadyInOption := sorted.options[newOption]
			if isAlreadyInOption {
				return Sorted{}, errors.New("found duplicate flag argument")
			}
			sorted.options[newOption] = []string{}
			sorted.argumentPos = append(sorted.argumentPos, job.Flag)
			currentOption = newOption
		} else {
			date, err := time.Parse(constants.ParseLayoutDate, arg)
			if err == nil && !date.IsZero() {
				sorted.dates = append(sorted.dates, date)
				sorted.argumentPos = append(sorted.argumentPos, job.Date)
			} else {
				sorted.options[currentOption] = append(sorted.options[currentOption], arg)
				sorted.argumentPos = append(sorted.argumentPos, job.Option)
			}
		}
	}
	return sorted, nil
}

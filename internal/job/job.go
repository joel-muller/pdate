package job

import (
	"errors"
	"pdate/internal/constants"
	"time"
)

type Argument int

const (
	Date Argument = iota
	Flag
	Option
)

type Language int

const (
	English Language = iota
	Spanish
	French
	German
	Swiss
	Italian
	Portuguese
	Dutch
	Russian
	Chinese
	Arabic
	Hindi
)

type Job struct {
	DatesInput      []time.Time
	PosArguments    []Argument
	IgnoredWeekdays []time.Weekday
	Reversed        bool
	Format          string
	Version         bool
	Help            bool
	Language        Language
}

func New() *Job {
	return &Job{
		[]time.Time{},
		[]Argument{},
		[]time.Weekday{},
		false,
		constants.DefaultInputFormat,
		false,
		false,
		English,
	}
}

func Validate(job *Job) error {
	if InvalidNumberOfDates(job) {
		return errors.New("wrong number of dates provided")
	}
	if DatesNextToEachOther(job) {
		return errors.New("the two dates are not next to each other")
	}
	if DatesBetweenOptions(job) {
		return errors.New("dates are in between options")
	}
	if DoubleWeekday(job) {
		return errors.New("double weekdays for ignore -i flag detected")
	}
	return nil
}

func InvalidNumberOfDates(job *Job) bool {
	if len(job.DatesInput) > 2 {
		return true
	}
	return false
}

func DatesNextToEachOther(job *Job) bool {
	dateIndex := -1
	for i, arg := range job.PosArguments {
		if arg == Date {
			if dateIndex != -1 {
				if i-1 != dateIndex {
					return true
				}
			} else {
				dateIndex = i
			}
		}
	}
	return false
}

func DatesBetweenOptions(job *Job) bool {
	for i, arg := range job.PosArguments {
		if arg == Date {
			if len(job.PosArguments) > i+1 && job.PosArguments[i+1] == Option {
				return true
			}
		}
	}
	return false
}

func DoubleWeekday(job *Job) bool {
	weekdays := make(map[time.Weekday]bool)
	for _, wd := range job.IgnoredWeekdays {
		if weekdays[wd] {
			return true
		} else {
			weekdays[wd] = true
		}
	}
	return false
}

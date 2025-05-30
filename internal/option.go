package internal

import (
	"errors"
)

type option int

const (
	IGNORE option = iota
	REVERSE
	DATE
)

var OptionKeyWords = map[string]option{
	"-i": IGNORE,
	"-r": REVERSE,
}

func SortArguments(args []string) (map[option][]string, error) {
	var options = map[option][]string{}
	var currentOption = DATE
	options[currentOption] = []string{}
	for _, arg := range args {
		if len(arg) == 0 {
			return map[option][]string{}, errors.New("empty argument provided")
		}
		if arg[0] == '-' {
			newOption, val := OptionKeyWords[arg]
			if !val {
				return map[option][]string{}, errors.New("found unknown argument")
			}
			_, isAlreadyInOption := options[newOption]
			if isAlreadyInOption {
				return map[option][]string{}, errors.New("found duplicate option argument")
			}
			options[newOption] = []string{}
			currentOption = newOption
		} else {
			options[currentOption] = append(options[currentOption], arg)
		}
	}
	return options, nil
}

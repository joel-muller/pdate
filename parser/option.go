package parser

type option int

const (
	IGNORE option = iota
)

var OptionKeyWords = map[option]string{
	IGNORE: "-i",
}

func IsOption(keyword string) bool {
	for _, value := range OptionKeyWords {
		if value == keyword {
			return true
		}
	}
	return false
}

func GetArgsForOptionInArgs(args []string, option option) []string {
	var index = -1
	for i, arg := range args {
		if arg == OptionKeyWords[option] {
			index = i
			break
		}
	}
	var optionArgs = []string{}
	if index == -1 {
		return optionArgs
	}
	for i := index + 1; i < len(args); i++ {
		if IsOption(args[i]) {
			return optionArgs
		}
		optionArgs = append(optionArgs, args[i])
	}
	return optionArgs
}

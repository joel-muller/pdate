package main

import (
	"fmt"
	"os"
	"pdate/internal"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if internal.NeedHelp(argsWithoutProg) {
		fmt.Println(internal.PrintHelp())
		return
	}
	dates, err := internal.GetDates(argsWithoutProg)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, date := range dates {
		fmt.Println(date)
	}
}

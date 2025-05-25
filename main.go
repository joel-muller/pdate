package main

import (
	"fmt"
	"os"
	"pdate/logic"
	"pdate/printer"
)

func main() {
	argsWithoutProg := os.Args[1:]
	dates, err := logic.GetDates(argsWithoutProg)
	if err != nil {
		fmt.Println("Error could not print dates")
	} else {
		printer.PrintDates(dates, "format")
	}
}

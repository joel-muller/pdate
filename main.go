package main

import (
	"fmt"
	"os"
	"pdate/internal/dates"
	"pdate/internal/job"
	"pdate/internal/parser"
)

func main() {
	argsWithoutProg := os.Args[1:]
	j := job.New()
	parseErr := parser.Parse(argsWithoutProg, j)
	if parseErr != nil {
		fmt.Println(parseErr)
		return
	}
	validErr := job.Validate(j)
	if validErr != nil {
		fmt.Println(validErr)
		return
	}
	for _, date := range dates.GetDates(j) {
		fmt.Println(date)
	}
}

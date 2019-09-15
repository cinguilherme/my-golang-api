package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("it begins")
}

func processEntry(entry string) (int, []string) {
	broken := s.Split(entry, ":")
	costumer := broken[0]
	dates := s.Split(s.Join(broken[1:], ""), ",")

	constumerType := 1
	if costumer == "regular" {
		constumerType = 2
	}

	return constumerType, dates
}

func buildSimplifiedWeekDayBoolArray(list []string) []bool {
	//zp := regexp.MustCompile(`(z+)`)
	var ret []bool

	for i := range list {

		if s.Contains("(sun)(sat)", list[i]) {
			ret = append(ret, true)
		} else {
			ret = append(ret, false)
		}
	}

	return ret
}

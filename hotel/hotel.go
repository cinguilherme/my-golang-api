package main

import (
	"fmt"
)

func main() {
	fmt.Println("it begins")
}

func processEntry(entry string) (int, []string) {
	dates := []string{"mon", "tue"}
	constumerType := 2
	return constumerType, dates
}

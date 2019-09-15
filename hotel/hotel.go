package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

type Hotel struct {
	Name       string
	WendFactor float32
	BasicTarif float32
}

type Hotels []Hotel

var arrHotels Hotels = Hotels{
	Hotel{Name: "woodlake", WendFactor: 1.3, BasicTarif: 25},
	Hotel{Name: "treelake", WendFactor: 1.1, BasicTarif: 22},
}

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

func getPriceForHotel(hotel Hotel, arrDAys []bool) float32 {
	var priceX float32 = 0.0
	for i := range arrDAys {
		wend := arrDAys[i]
		price := hotel.BasicTarif
		if wend == true {
			price = price * hotel.WendFactor
		}
		priceX += price
	}
	return priceX
}

func getCheapest(hotels Hotels, arrDays []bool) Hotel {
	var cheapest Hotel = hotels[0]
	for i := range hotels {
		getPriceForHotel(hotels[i], arrDays)
	}

	return cheapest
}

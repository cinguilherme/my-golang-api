package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestProcessEntry(t *testing.T) {
	p("Hotel test class")
	tp, arr := processEntry("regular:(mon),(tue),(sat)")

	assert.Assert(t, tp == 2)
	assert.Assert(t, len(arr) == 3)
}

func TestBuildSimplifiedWeekDayBoolArray(t *testing.T) {
	var boolArray []bool
	input := []string{"(sun)", "(mon)"}
	boolArray = buildSimplifiedWeekDayBoolArray(input)

	assert.Assert(t, len(boolArray) == 2)
	assert.DeepEqual(t, boolArray, []bool{true, false})
}

func TestGetPriceForHotel(t *testing.T) {
	var hotel Hotel = Hotel{
		Name: "testHotel", WendFactor: 1.5, BasicTarif: 10.0,
	}
	var price float32 = getPriceForHotel(
		hotel, []bool{true, true, false})
	assert.Assert(t, price == 40.0)
}

func TestGetPriceForHotelOther(t *testing.T) {
	var hotel Hotel = Hotel{
		Name: "testHotel", WendFactor: 1.1, BasicTarif: 10.0,
	}
	var price float32 = getPriceForHotel(
		hotel, []bool{false, false})
	var expected float32 = 20.0
	assert.Equal(t, price, expected)
}

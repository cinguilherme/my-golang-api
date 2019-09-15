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

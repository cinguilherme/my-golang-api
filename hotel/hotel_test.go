package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestEntry(t *testing.T) {
	p("Hotel test class")
	tp, arr := processEntry("regular:(mon),(tue),(sat)")

	assert.Assert(t, tp == 2)
	assert.Assert(t, len(arr) == 3)
}

package main

import (
	"fmt"
	"testing"
)

func TestEntry(t *testing.T) {
	tp, arr := processEntry("regular:(mon),(tue),(sat)")
	fmt.Println(tp, arr)

}

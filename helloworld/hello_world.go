package main

import "fmt"

var globalOne = 10

var rawString string = `

	This is me just messing everything up really bad..

`

func main() {
	numBytes, _ := fmt.Println("hi")
	fmt.Println(numBytes)

	foo()

	fmt.Println("something else here")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}

	}
	otherFoo(10, 20, 40)
}

func otherFoo(val ...int) {
	fmt.Println(val, inter(),
		stringer(), loko(true), rawString, buildMyArrayOfStrings())
}

func inter() int {
	var prefefined int = 10
	var predefNoVal int
	res := 10 + 100 + prefefined
	predefNoVal = 15 + globalOne
	res += predefNoVal
	return res
}

func loko(opt bool) interface{} {
	return "that????"
}

func stringer() string {
	res := "base"
	res += " how this responds?"
	return res
}

func foo() {
	fmt.Println("I`m foo")
}

func buildMyArrayOfStrings() []string {
	var some []string = []string{"string", "string2"}
	return some
}

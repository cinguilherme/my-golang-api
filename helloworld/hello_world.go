package main

import "fmt"

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
	fmt.Println(val)
}

func foo() {
	fmt.Println("I`m foo")
}

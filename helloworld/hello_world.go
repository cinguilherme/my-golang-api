package main

import "fmt"

func main() {
	fmt.Printf("hi")
	foo()

	fmt.Println("something else here")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}

	}
}

func build_my_string() {
	str := "here is my string beeing built"
	return str
}

func foo() {
	fmt.Println("I`m foo")
}
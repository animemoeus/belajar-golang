package main

import "fmt"

func deferFunction() {
	fmt.Println("Defer function")
}

func panicFunction() {
	panic("Ups, ada error")
}

func sayHello(error bool) {
	defer deferFunction()

	if error {
		panicFunction()
	}
}

func main() {
	sayHello(true)
}

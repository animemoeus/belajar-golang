package main

import "fmt"

func end() {
	fmt.Println("Defer function")
}

func sayHello(name string) {
	defer end()

	fmt.Println("Hello", name)
}

func main() {
	sayHello("arter")
}

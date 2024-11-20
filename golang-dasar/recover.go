package main

import "fmt"

func deferFunction() {
	fmt.Println("End App")

	panicMessage := recover()

	if panicMessage != nil {
		fmt.Println("Terjadi panic:", panicMessage)
	}
}

func panicFunction() {
	defer deferFunction()
	panic("Ups error...")

}

func main() {
	//defer deferFunction()
	panicFunction()
}

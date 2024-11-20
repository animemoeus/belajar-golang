package main

import "fmt"

type Filter func(string) string // type declaration

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}

	return name
}

func main() {
	sayHelloWithFilter("anjing", spamFilter)
}

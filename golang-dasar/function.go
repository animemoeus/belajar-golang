package main

import "fmt"

func sayHello() {
	fmt.Println("halo dunia")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func main() {
	//sayHello()
	//sayHelloTo("arter", "tendean")
	hello := getHello("arter tendean")
	fmt.Println(hello)

	firstName, lastName := getFullName("arter", "tendean")
	fmt.Println(firstName, lastName)
}

package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, my name is", customer.Name)
}

func (customer Customer) callSomeone(crush Customer) {
	fmt.Println("My name is", customer.Name)
	crush.sayHello()
}

func main() {
	var arter Customer
	arter.Name = "Arter Tendean"
	arter.Address = "Tokyo"
	arter.Age = 100

	fmt.Println(arter)

	angie := Customer{
		Name: "ahsl",
		Age:  19,
	}

	fmt.Println(angie)

	tkg := Customer{"Takagi", "Tokyo", 19}
	tkg.sayHello()

	arter.callSomeone(angie)
}

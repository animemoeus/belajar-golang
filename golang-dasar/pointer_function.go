package main

import "fmt"

type Address struct {
	Country string
}

func updateAddress(address *Address) {
	address.Country = "Jepang"
}

func main() {
	var x Address = Address{"Indonesia"}
	updateAddress(&x)

	fmt.Println(x)
}

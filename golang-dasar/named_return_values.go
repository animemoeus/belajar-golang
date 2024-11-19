package main

import "fmt"

func getCompletename() (firstName string, lastName string) {
	firstName = "arter"
	lastName = "tendean"
	return firstName, lastName
}

func main() {
	firstName, lastName := getCompletename()

	fmt.Println(firstName, lastName)
}

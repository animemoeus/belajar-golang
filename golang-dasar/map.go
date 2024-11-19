package main

import "fmt"

func main() {
	person := map[string]string{
		"firstName": "Arter",
		"lastName":  "Tendean",
		"toDelete":  "ahsl",
	}

	//fmt.Println("person:", person)
	fmt.Println("First Name: ", person["firstName"])
	fmt.Println("Last Name: ", person["lastName"])
	fmt.Println("Unknown: ", person["hahaha"])

	fmt.Println("Person Length: ", len(person))
	fmt.Println("person:", person)
	delete(person, "toDelete")
	fmt.Println("Person Length: ", len(person))
	fmt.Println("person:", person)
}

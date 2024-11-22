package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	//arter := Address{"Manado", "Sulawesi Utara", "Indonesia"}
	//tendean := arter
	//
	//tendean.Country = "Jepang"
	//
	//fmt.Println(arter)
	//fmt.Println(tendean)

	arter := Address{"Manado", "Sulawesi Utara", "Indonesia"}
	var tendean *Address = &arter

	tendean.Country = "Jepang"

	fmt.Println(arter)
	fmt.Println(tendean)

}

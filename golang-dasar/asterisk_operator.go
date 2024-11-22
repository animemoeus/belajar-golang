package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	arter := Address{"Manado", "Sulawesi Utara", "Indonesia"}
	tendean := &arter

	tendean.Country = "Jepang"

	fmt.Println(arter)
	fmt.Println(tendean)

	*tendean = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(arter)
	fmt.Println(tendean)

	alamat1 := new(Address)
	var alamat2 *Address = alamat1

	alamat1.Province = "Province"
	alamat2.City = "City"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}

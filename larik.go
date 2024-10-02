package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Alpha"
	names[1] = "Beta"
	names[2] = "Gamma"

	names[1] = "Delta"

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//	create array directly
	var values = [...]int{1, 3}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	//fmt.Println(values[2])

}

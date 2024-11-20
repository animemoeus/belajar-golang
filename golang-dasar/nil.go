package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	x := NewMap("arter")

	if x == nil {
		fmt.Println("kosong")
	} else {
		fmt.Println(x)
	}

}

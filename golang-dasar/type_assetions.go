package main

import "fmt"

func random() any {
	return "ok"
}

func main() {
	x := random()
	y := x.(string)
	fmt.Println(y)

}

package main

import "fmt"

func sayGoodBye(name string) string {
	return "Good Bye" + name
}

func main() {
	goodBye := sayGoodBye
	fmt.Println(goodBye("arter"))
}

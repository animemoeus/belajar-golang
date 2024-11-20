package main

import "fmt"

func main() {

	switch name := "tendean"; name {
	case "arter":
		fmt.Println("halo arter")
	case "tendean":
		fmt.Println("halo tendean")
	default:
		fmt.Println("kamu siapa?")
	}

}

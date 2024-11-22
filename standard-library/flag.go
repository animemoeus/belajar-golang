package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "default username", "input your username")

	flag.Parse()
	fmt.Println(*username)
}

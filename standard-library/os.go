package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for i, arg := range args {
		fmt.Println(i+1, arg)
	}

}

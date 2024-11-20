package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) createWeebs() {
	man.Name += "aowkaowk"
}

func main() {
	arter := Man{"Arter Tendean"}
	arter.createWeebs()

	fmt.Println(arter)
}

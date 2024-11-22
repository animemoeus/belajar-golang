package main

import (
	"errors"
	"fmt"
)

var ValidationError = errors.New("Validation Error")
var UnknownError = errors.New("Unknown Error")

func createPost(title string) (string, error) {
	if title == "" {
		return "", ValidationError
	}

	if len(title) > 10 {
		return "", UnknownError
	}

	return title, nil
}

func main() {

	//post, err := createPost("")
	var post string
	var err error
	post, err = createPost("arter tendean")

	if err != nil {
		//fmt.Println("Error:", err)
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error hahaha")
		} else {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println(post)

}

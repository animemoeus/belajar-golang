package main

import (
	"fmt"
	"io"
)
import "net/http"

func main() {
	fmt.Println("Hello, arterrrr")
	resp, err := http.Get("https://api.animemoe.us/waifu/")

	if err != nil {
		fmt.Println("ada kesalahan")
	}

	if resp != nil {
		fmt.Println("berhasil, berhasil, berhasil, horeee")
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Body)
	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))
	fmt.Println(err)
}

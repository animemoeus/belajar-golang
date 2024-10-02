package main

import "fmt"
import "net/http"

func main() {
	fmt.Println("Hello, arterrrr")
	resp, err := http.Get("https://api.animemoe.us/waifu/random/")
	
	if err != nil {
		fmt.Println("ada kesalahan")
	}

	if resp != nil {
		fmt.Println("berhasil, berhasil, berhasil, horeee")
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Body)
}

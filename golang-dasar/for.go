package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println("counter: ", counter)
	//	counter++
	//}

	//for i := 0; i <= 10; i++ {
	//	for j := 0; j < i; j++ {
	//		fmt.Print("*")
	//	}
	//	fmt.Println()
	//}

	waifu := []string{"Takagi", "Kaguya", "Angie"}
	fmt.Println("waifu: ", waifu)

	for i := 0; i < len(waifu); i++ {
		fmt.Println(waifu[i])
	}

	for index, value := range waifu {
		fmt.Println(index, value)
	}
}

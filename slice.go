package main

import "fmt"

func main() {
	//names := [...]string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet"}

	// membuat slice dari array
	//var slice1 []string = names[1:]
	//fmt.Println(slice1)
	////
	////	slice function
	//fmt.Println("panjang slice1:", len(slice1))
	//slice2 := append(slice1, "arter tendean")
	//fmt.Println("slice2: ", slice2)
	//
	//days := [...]string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	//fmt.Println("days: ", days)
	//
	//var daySlice1 []string = days[5:]
	//fmt.Println("daySlice1: ", daySlice1)
	//daySlice1[0] = "Minggu Baru"
	//fmt.Println("days: ", days)
	//
	//daySlice2 := append(daySlice1, "Hari Libur")
	//daySlice2[0] = "Ups"
	//fmt.Println("daySlice2: ", daySlice2)
	//fmt.Println("days: ", days)
	//
	//newSlice := make([]string, 2, 5)
	//newSlice[0] = "arter"
	//newSlice[1] = "tendean"
	//fmt.Println("newSlice: ", newSlice)
	//
	//newSlice2 := append(newSlice, "ahsl")
	//fmt.Println("newSlice2: ", newSlice2)
	//newSlice2[0] = "Updated"
	//fmt.Println("newSlice: ", newSlice)
	//fmt.Println("newSlice2: ", newSlice2)

	var days = []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}
	fmt.Println("days: ", days)

	daySlice1 := days[4:5]
	daySlice1[0] = "Jumat Baru"
	fmt.Println("daySlice1: ", daySlice1)
	fmt.Println("days: ", days)

	daySlice2 := append(daySlice1, "Hari Libur")
	fmt.Println("daySlice2: ", daySlice2)
	fmt.Println("days: ", days)
}

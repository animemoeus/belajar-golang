package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	utc := time.Date(2005, time.February, 15, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	//time.RFC822
	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Failed to format time:", err)
	}

	fmt.Println("Formated Time:", valueTime)
}

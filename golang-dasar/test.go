package main

import "fmt"

func main() {
	a := 10 // Ini variabel biasa
	b := &a // Ini pointer, alias alamat memori si 'a'

	fmt.Println("Nilai a:", a)            // 10
	fmt.Println("Alamat a:", b)           // Misalnya 0xc0000140a0
	fmt.Println("Nilai di alamat b:", *b) // 10

	*b = 20                         // Ubah nilai 'a' lewat pointer
	fmt.Println("Nilai baru a:", a) // 20
}

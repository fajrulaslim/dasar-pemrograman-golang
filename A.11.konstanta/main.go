package main

import "fmt"

func main() {
	fmt.Println("KONSTANTA")
	fmt.Println("---------")

	// PENGGUNAAN KONSTANTA
	const firstName string = "John"
	const lastName = "Wick"
	fmt.Print("Halo ", firstName, " ", lastName, "\n")

	// PENGGUNAAN FUNGSI PRINT
	fmt.Println("---------")
	fmt.Println("John Wick")
	fmt.Println(firstName, lastName)

	fmt.Print("John Wick \n")
	fmt.Print("John ", "Wick \n")
	fmt.Print("John", " ", "Wick \n")

	// DEKLARASI MULTIKONSTANTA
	fmt.Println("---------")
	const (
		square         = "Kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.4
	)
	fmt.Println(square, isToday, numeric, floatNum)

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"
	fmt.Println(satu, dua)
	fmt.Println(three, four)
}

package main

import "fmt"

func main() {
	fmt.Println("TIPE DATA")
	fmt.Println("--------")

	// TIPE DATA NUMERIK NON-DESIMAL
	fmt.Println("--------")
	var positiveNumber uint8 = 89
	var negativeNumber int32 = -1243423644

	fmt.Printf("bilangan positif: %d \n", positiveNumber)
	fmt.Printf("bilangan negatif: %d \n", negativeNumber)

	// TIPE DATA NUMERIK DESIMAL
	fmt.Println("--------")
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f \n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f \n", decimalNumber)

	// TIPE DATA BOOLEAN
	fmt.Println("--------")
	var exist bool = true
	fmt.Printf("exist: %t \n", exist)

	// TIPE DATA STRING
	fmt.Println("--------")
	var message string = "Pesan"
	fmt.Println(message)

	// NILAI NIL DAN ZERO VALUE
	fmt.Println("--------")
}

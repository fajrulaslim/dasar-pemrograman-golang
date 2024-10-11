package main

import "fmt"

func main() {
	fmt.Println("VARIABEL")
	fmt.Println("--------")

	// DEKLARASI VARIABEL BESERTA TIPE DATA
	var firstName string = "John"
	var lastName string = "Wick"

	fmt.Println(firstName, lastName)

	// FUNGSI PRINTF
	fmt.Println("--------")
	fmt.Printf("Halo John Wick! \n")
	fmt.Printf("Halo %s %s! \n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	// DEKLARASI VARIABEL TANPA TIPE DATA
	fmt.Println("--------")
	middeName := "D."
	fmt.Printf("Halo %s %s %s! \n", firstName, middeName, lastName)

	// DEKLARASI MULTIVARIABEL
	fmt.Println("--------")
	var first, second, third string = "satu", "dua", "tiga"
	enam, tujuh, delapan := "enam", "tujuh", "delapan"
	fmt.Println(first, second, third)
	fmt.Println(enam, tujuh, delapan)

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one, isFriday, twoPointTwo, say)

	// VARIABEL UNDERSCORE
	fmt.Println("--------")
	_ = "belajar golang"
	_ = "dasar"
	title, _ := "tutorial", "belajar"
	fmt.Println(title)

	// DEKLARASI VARIABEL DENGAN NEW
	fmt.Println("--------")
	fruit := new(string)
	fmt.Println(fruit)
	fmt.Println(*fruit)
}

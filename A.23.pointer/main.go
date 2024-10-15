package main

import "fmt"

func main() {
	fmt.Println("POINTER")

	// PENERAPAN
	fmt.Println("-------")
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) \t:", numberA)
	fmt.Println("numberA (address) \t:", &numberA)

	fmt.Println("numberB (value) \t:", *numberB)
	fmt.Println("numberB (address) \t:", numberB)

	// EFEK PERUBAHAN NILAI POINTER
	fmt.Println("-------")
	numberA = 5

	fmt.Println("numberA (value) \t:", numberA)
	fmt.Println("numberA (address) \t:", &numberA)
	fmt.Println("numberB (value) \t:", *numberB)
	fmt.Println("numberB (address) \t:", numberB)

	// PARAMETER POINTER
	fmt.Println("-------")
	var number = 4
	fmt.Println("before:", number)

	change(&number, 10)
	fmt.Println("after:", number)
}

func change(original *int, value int) {
	*original = value
}

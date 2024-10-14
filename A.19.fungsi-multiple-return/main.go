package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("FUNGSI MULTIPLE RETURN")

	// PENERAPAN
	fmt.Println("----------------------")
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Printf("Luas \t\t : %.2f \n", area)
	fmt.Printf("Keliling \t : %.2f \n", circumference)

	// PREDEFINED RETURN VALUE
	fmt.Println("----------------------")
	var diameter2 float64 = 14
	var area2, circumference2 = calculate2(diameter2)
	fmt.Printf("Luas \t\t : %.2f \n", area2)
	fmt.Printf("Keliling \t : %.2f \n", circumference2)
}

func calculate(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2) // pow untuk pangkat
	var circumference = math.Pi * d

	return area, circumference
}

func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}

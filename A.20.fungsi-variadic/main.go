package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("FUNGSI VARIADIC")

	// PENERAPAN
	fmt.Println("---------------")
	var avg = calculate(2, 3, 4, 5, 6, 7, 8, 9)
	var msg = fmt.Sprintf("Rata-rata : %.2f\n", avg)
	fmt.Println(msg)

	// FUNGSI VARIADIC MENGGUNAKANA SLICE
	fmt.Println("---------------")
	var numbers = []int{2, 3, 4, 5, 6, 7, 8, 9}
	var avg2 = calculate(numbers...)
	var msg2 = fmt.Sprintf("Rata-rata : %.2f\n", avg2)
	fmt.Println(msg2)

	// FUNGSI PARAMETER BIASA DAN VARIADIC
	fmt.Println("---------------")
	yourHobbies("Wick", "running", "swimming")
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, " ")

	fmt.Printf("Hello, my name is %s \n", name)
	fmt.Printf("My hobbies are %s\n", hobbiesAsString)
}

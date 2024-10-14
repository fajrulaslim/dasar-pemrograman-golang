package main

import "fmt"

func main() {
	fmt.Println("CLOSURE")

	// CLOSURE DISIMPAN DALAM VARIABEL
	fmt.Println("--------")
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				min, max = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers1 = []int{2, 3, 4, 5, 6, 7, 8, 9}
	var min, max = getMinMax(numbers1)
	fmt.Printf("data: %v \n min: %v \n max: %v", numbers1, min, max)

	// IMMEDIETELY INVOKED FUNCTION EXPRESSIONS (IIFE)
	fmt.Println("--------")
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)

	// CLOSURE SEBAGAI NILAI KEMBALIAN
	fmt.Println("--------")
	var maxNumber = 3
	var howMany, getNumbers = findMax(numbers, maxNumber)
	var theNumbers = getNumbers()
	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e < max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}

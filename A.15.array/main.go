package main

import "fmt"

func main() {
	fmt.Println("ARRAY")
	fmt.Println("-----")

	var names [4]string
	names[0] = "Jhon"
	names[1] = "Wick"
	fmt.Println(names[0], names[1])

	// INISIALISASI NILAI AWAL ARRAY
	fmt.Println("-----")
	var fruits = [3]string{"apple", "mango", "papay"}
	fmt.Println("Jumlah element \t :", len(fruits))
	fmt.Println("Isi element \t :", fruits)

	var animals = [3]string{
		"Kucing",
		"Anjing",
		"Lalat",
	}
	fmt.Println(animals)

	// INISIALISASI ARRAY TANPA JUMLAH ELEMENT
	fmt.Println("-----")
	var numbers = [...]int{1, 2, 3}
	fmt.Println(numbers)

	// ARRAY MULTIDIMENSI
	fmt.Println("-----")
	var numbers1 = [2][3]int{{2, 3, 4}, {1, 2, 3}}
	fmt.Println(numbers1)

	// PERULANGAN ARRAY DENGAN FOR
	fmt.Println("-----")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("element %d : %s \n", i, fruits[i])
	}

	// PERUALANGAN ARRAY DENGAN FOR-RANGE
	fmt.Println("-----")
	for i, fruit := range fruits {
		fmt.Printf("element %d : %s \n", i, fruit)
	}

	// PENGGUNAAN UNDERSCORE DALAM FOR-RANGE
	fmt.Println("-----")
	for _, fruit := range fruits {
		fmt.Printf("element : %s \n", fruit)
	}

	// ALOKASI ELEMENT ARRAY DENGAN MAKE
	var fruits1 = make([]string, 2)
	fmt.Println("-----")
	fruits1[0] = "banana"
	fruits1[1] = "apple"
	fmt.Println(fruits1)
}

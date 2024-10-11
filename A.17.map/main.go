package main

import "fmt"

func main() {
	fmt.Println("MAP")
	fmt.Println("---")

	var chicken = map[string]int{}

	chicken["januari"] = 1
	chicken["februari"] = 2

	fmt.Println("Januari: ", chicken["januari"])
	fmt.Println("Februari: ", chicken["februari"])

	fmt.Println("---")
	var data1 map[string]int
	fmt.Println(data1)
	var data2 = map[string]int{}
	fmt.Println(data2)

	fmt.Println("---")
	var chicken1 = map[string]int{"januari": 1, "februari": 2}
	var chicken2 = map[string]int{
		"januari":  1,
		"februari": 2,
	}
	fmt.Println(chicken1)
	fmt.Println(chicken2)
}

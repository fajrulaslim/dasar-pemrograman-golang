package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("FUNGSI DI DALAM PARAMETER")

	// PENERAPAN
	fmt.Println("-------------------------")
	var data = []string{"John", "Wick", "Ethan"}

	var dataConstain0 = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("Data asli:", data)
	fmt.Println("filter ada huruf \"o\"\t:", dataConstain0)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)

	// ALIAS SKEMA CLOSURE
	fmt.Println("-------------------------")
	var dataConstain1 = filter2(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength6 = filter2(data, func(each string) bool {
		return len(each) == 6
	})

	fmt.Println("Data asli:", data)
	fmt.Println("filter ada huruf \"o\"\t:", dataConstain1)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength6)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

type FilterCallback func(string) bool

func filter2(data []string, callback FilterCallback) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

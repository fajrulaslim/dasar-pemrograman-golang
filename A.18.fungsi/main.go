package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource((time.Now().Unix())))

func main() {
	fmt.Println("FUNGSI")
	fmt.Println("------")

	// PENERAPAN FUNGSI
	var names = []string{"John", "Wick"}
	printMessage("Halo", names)

	// FUNGSI DENGAN RETURN VALUE
	fmt.Println("------")
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println(randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println(randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println(randomValue)

	// KEYWORD RETURN UNTUK MENGHENTIKAN PROSES
	fmt.Println("------")
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)

}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %d \n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d \n", m, n, res)
}

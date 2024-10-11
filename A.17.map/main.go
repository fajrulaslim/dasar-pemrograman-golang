package main

import "fmt"

func main() {
	fmt.Println("MAP")
	fmt.Println("---")

	// PENGGUNAAN MAP
	var chicken = map[string]int{}

	chicken["januari"] = 1
	chicken["februari"] = 2
	chicken["maret"] = 3
	chicken["april"] = 4
	chicken["mei"] = 5

	fmt.Println("Januari: ", chicken["januari"])
	fmt.Println("Februari: ", chicken["februari"])

	// INILIASISASI NILAI MAP
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

	fmt.Println("---")
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)

	// ITERASI ITEM MAP DENGAN FOR-RANGE
	fmt.Println("---")
	for key, val := range chicken {
		fmt.Println(key, "\t : ", val)
	}

	// MENGHAPUS ITEM MAP
	fmt.Println("---")
	delete(chicken, "januari")
	fmt.Println(chicken)

	// DETEKSI KEBERADAAN ITEM DENGAN KEY TERTENTU
	fmt.Println("---")
	var value, isExist = chicken["maret"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Tidak ada")
	}

	// KOMBINASI MAP DAN SLICE
	fmt.Println("---")
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	var data = []map[string]string{
		{"name": "chicken red", "gender": "male", "color": "red"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	for _, dt := range data {
		fmt.Println(dt)
	}

}

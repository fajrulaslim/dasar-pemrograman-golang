package main

import "fmt"

func main() {
	fmt.Println("SLICE")
	fmt.Println("-----")

	// INISIALISASI SLICE
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits)
	fmt.Println(fruits[0])

	// HUBUNGAN SLICE DENGAN ARRAY & OPERASI SLICE
	fmt.Println(fruits[0:2])
	fmt.Println(fruits[0:3])
	fmt.Println(fruits[0:0])
	fmt.Println(fruits[3:3])
	fmt.Println(fruits[:])
	fmt.Println(fruits[0:])
	fmt.Println(fruits[:3])

	fmt.Println("-----")
	var arrFruits = [3]string{"apple", "grape", "banana"}
	fmt.Println(arrFruits) // array
	var newSliceArrFruits = arrFruits[0:1]
	fmt.Println(newSliceArrFruits)

	// SLICE SEBAGAI TIPE DATA REFERENCE
	fmt.Println("-----")
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var bbFruits = bFruits[0:1]

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	bbFruits[0] = "pineaple"
	fmt.Println("")
	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	// FUNGSI LEN DAN CAP
	fmt.Println("-----")
	fmt.Println("len fruits", len(fruits))
	fmt.Println("cap fruits", cap(fruits))
	fmt.Println("len aFruits", len(aFruits))
	fmt.Println("cap aFruits", cap(aFruits))
	fmt.Println("len bFruits", len(bFruits))
	fmt.Println("cap bFruits", cap(bFruits))

	// FUNGSI APPEND
	fmt.Println("-----")
	var cFruits = append(fruits, "papaya")
	fmt.Println(fruits)
	fmt.Println(cFruits)
	var dFruits = append(aFruits, "papaya")
	fmt.Println(aFruits)
	fmt.Println(dFruits)
	var eFruits = append(bFruits, "papaya")
	fmt.Println(bFruits)
	fmt.Println(eFruits)

	// FUNGSI COPY
	fmt.Println("-----")
	dst := make([]string, 6)
	src := []string{"apple", "banana", "watermelon", "papaya"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)

	// AKSES SLICE DENGAN 3 INDEKS
	fmt.Println("-----")
	var nFruits = fruits[0:2]
	var mFruits = fruits[0:2:4]

	fmt.Println(nFruits)
	fmt.Println(mFruits)
	fmt.Println(len(mFruits))
	fmt.Println(cap(mFruits))
}

package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	// s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	fmt.Println("METHOD")

	// PENERAPAN
	fmt.Println("------")
	var s1 = student{"John Wick", 24}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan:", name)

	// METHOD POINTER
	fmt.Println("------")
	s1.changeName1("jason bourne")
	fmt.Println("s1 after changeName1", s1.name)

	s1.changeName2("ethan hunt")
	fmt.Println("s2 after changeName2", s1.name)

	var s2 = &student{"ahmad kahfi", 22}
	s2.sayHello()
}

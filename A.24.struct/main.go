package main

import "fmt"

type student struct {
	name  string
	grade int
}

type person struct {
	name string
	age  int
}

type teacher struct {
	grade int
	person
}

type manager struct {
	person
	age   int
	grade int
}

func main() {
	fmt.Println("STRUCT")

	// DEKLARASI STRUCT
	fmt.Println("------")
	var s1 student
	s1.name = "John Wick"
	s1.grade = 2

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)
	fmt.Println(s1)

	// INISIALISASI OBJECT STRUCT
	fmt.Println("------")
	var s2 = student{"ethan", 2}
	var s3 = student{name: "jason"}

	fmt.Println("student :", s2.name)
	fmt.Println(s2)
	fmt.Println("student :", s3.name)
	fmt.Println(s3)

	// VARIABEL OBJECT POINTER
	fmt.Println("------")
	var student1 = student{name: "Wick"}

	var student2 *student = &student1
	fmt.Println("student 1, name:", student1.name)
	fmt.Println("student 2, name:", student2.name)

	student2.name = "Ethan"
	fmt.Println("student 1, name", student1.name)
	fmt.Println("student 2, name", student2.name)

	// EMBEDDED STRUCT
	fmt.Println("------")
	var t1 = teacher{}
	t1.name = "Wick"
	t1.age = 21
	t1.grade = 2

	fmt.Println("name :", t1.name)
	fmt.Println("age :", t1.age)
	fmt.Println("age :", t1.person.age)
	fmt.Println("grade :", t1.grade)

	// EMBEDDED STRUCT DENGAN NAMA PROPERTY SAMA
	fmt.Println("------")
	var m1 = manager{}
	m1.name = "Wick"
	m1.age = 21        // age of manager
	m1.person.age = 22 // age or person
	m1.grade = 9

	fmt.Println(m1.name)
	fmt.Println(m1.age)
	fmt.Println(m1.person.age)
	fmt.Println(m1.grade)

	// PENGISIAN NILAI SUB-STRUCT
	fmt.Println("------")
	var p1 = person{name: "Wick", age: 21}
	var q1 = teacher{person: p1, grade: 2}

	fmt.Println(q1.name)
	fmt.Println(q1.age)
	fmt.Println(q1.grade)

	// ANONYMOUS STRUCT
	fmt.Println("------")
	var a1 = struct {
		person
		grade int
	}{} // anonymous struct tanpa pengisian property
	a1.person = person{"Wick", 22}
	a1.grade = 2

	fmt.Println("name:", a1.person.name)
	fmt.Println("age:", a1.person.age)
	fmt.Println("grade:", a1.grade)

	var a2 = struct {
		person
		grade int
	}{
		person: person{"John", 21},
		grade:  3,
	} // anonymous strut dengan pengisian property
	fmt.Println("name:", a2.person.name)
	fmt.Println("age:", a2.person.age)
	fmt.Println("grade:", a2.grade)

	// KOMBINASI SLICE DAN STRUCT
	fmt.Println("------")
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 24},
		{name: "Bourne", age: 22},
	}
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	// INISIALISASI SLICE ANONYMOUS STRUCT
	fmt.Println("------")
	var allTeachers = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 1},
		{person: person{"john", 24}, grade: 2},
	}
	for _, teacher := range allTeachers {
		fmt.Println(teacher)
	}

	// DEKLARASI ANONYMOUS STRUCT DENGAN KEYWORD VAR
	fmt.Println("------")
}

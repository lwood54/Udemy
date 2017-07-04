package main

import "fmt"

// Person defines a person
type Person struct {
	First string
	Last  string
	Age   int
}

// DoubleZero defines agent characteristics
type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}
	fmt.Println(p1)
}

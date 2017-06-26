package main

import (
	"fmt"
)

func main() {
	a := 43
	c := "some string"

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	var d *string = &c

	fmt.Println(b)
	fmt.Println(d)

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer" --- so it points to a location in memory that stores an int
	// *int -- the * is part of the type -- b is of type *int
}

package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	// this shorthand can only be used from WITHIN a func()
	// the types are infered with this shorthand

	// %v uses the "default" for whatever type the variable is
	// %T expresses the TYPE for the variable
	fmt.Printf("%v is of Type: %T \n", a, a)
	fmt.Printf("%v is of Type: %T \n", b, b)
	fmt.Printf("%v is of Type: %T \n", c, c)
	fmt.Printf("%v is of Type: %T \n", d, d)

	// initializing but not assigning results in go assigning a zero value
	// which changes depending on the type.
	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v is of Type: %T \n", e, e)
	fmt.Printf("%v is of Type: %T \n", f, f)
	fmt.Printf("%v is of Type: %T \n", g, g)
	fmt.Printf("%v is of Type: %T \n", h, h)
}

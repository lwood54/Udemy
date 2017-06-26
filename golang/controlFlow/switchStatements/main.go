package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Hello Daniel")
	case "Medhi":
		fmt.Println("Hello Medhi")
	case "Jenny":
		fmt.Println("Hello Jenny")
	default:
		fmt.Println("Have you no friends?")
	}

	// Example: Explicitly stating fallthrough
	switch "Marcus" {
	case "Tim":
		fmt.Println("Hi Tim")
	case "Jenny":
		fmt.Println("Hi Jenny")
	case "Marcus":
		fmt.Println("Hi Marcus") // this will print
		fallthrough              // fallthrough tells it to move to the next line and execute
	case "Medhi":
		fmt.Println("Hi Medhi") // this will ALSO print
		fallthrough             // keeps going...
	case "Julian":
		fmt.Println("Hi Julian") // this will also print
	case "Sushant":
		fmt.Println("Hi Sushant") // no fallthrough, so this won't print
	}

	// EXAMPLE: multiple cases with the same output:
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Hello Tim...or Jenny.")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with 'M' ")
	case "Julian", "Sushant":
		fmt.Println("Hello Julian / Sushant")
	}

	// You can also have no switch statement, then it will go through each
	// case and identify the FIRST ONE that is true

	// You can also switch on TYPE
	// normally you switch on value of variable
	// go allows you to switch on TYPE of variable

	SwitchOnType(55)
}

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

/*
  - no default fallthrough
  - fallthrough is optional
  --- you can specify fallthrough by explicitly stating it
  --- break isn't needed like in other languages
*/

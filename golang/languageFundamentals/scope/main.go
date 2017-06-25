package main

import "fmt"

// package level scope:
// this even includes if you have multiple files within the same package,
// the variable would be available to use (or mess up) throughout the
// whole package
var x int = 42

var y = 0

func increment() int {
	y++
	return y
}

func main() {
	fmt.Println(x)
	foo()
	fmt.Println("1st call of increment(): ", increment())
	fmt.Println("2nd call of increment(): ", increment())
}

func foo() {
	fmt.Println(x)
}

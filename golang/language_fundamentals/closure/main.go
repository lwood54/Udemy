package main

import "fmt"

func main() {
	x := 0
	// anonymous function being defined to the variable "increment"
	increment := func() int {
		x++
		return x
	}
	fmt.Println("1st call of increment(): ", increment())
	fmt.Println("2nd call of increment(): ", increment())
}

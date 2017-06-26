package main

import "fmt"

// a function called wrapper expects to have a return value that is a func of type int
func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

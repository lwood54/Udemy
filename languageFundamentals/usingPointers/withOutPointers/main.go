package main

import "fmt"

func zero(z int) int {
	z = 0
	return z
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
}

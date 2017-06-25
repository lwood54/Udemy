package main

import "fmt"

// const p string = "death & taxes"
const q int = 42
const m = 55

const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

func main() {
	// const q = 42
	const p = "death & taxes"
	m := 88 // not quite sure why this is allowed to change if a const is an "unchanging value"
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	fmt.Println("m - ", m)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

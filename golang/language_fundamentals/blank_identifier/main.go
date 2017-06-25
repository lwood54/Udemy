package main

import "fmt"

func main() {
	a := "stored in a"
	// b := "stored in b"
	fmt.Println("a - ", a)
	// b is not being used - invalid code
	// a blank identifier tells Go that you know there is nothing there,
	// but you want to continue any way.
}

package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address in standard hex form: ", &a)
	fmt.Printf("a's memory address in decimal form: %d \n", &a)
}

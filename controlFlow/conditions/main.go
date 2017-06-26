package main

import "fmt"

func main() {

	// standard for loop:
	for i := 0; i < 10; i++ {
		fmt.Println("Hello: ", i)
	}

	// similar to a while statement
	j := 0
	for j < 10 {
		fmt.Println("Hi there: ", j)
		j++
	}

	// how to break out of a loop, even an infinite loop
	k := 0
	for {
		fmt.Println("Hola mi amigo: ", k)
		if k >= 9 {
			break
		}
		k++
	}

	m := 0
	for {
		m++
		if m > 50 {
			break
		} else if m%2 == 0 {
			fmt.Println("Even number:   ", m)
			continue
		} else {
			fmt.Println("Odd number: ", m)
		}
	}
}

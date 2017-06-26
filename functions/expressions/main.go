package main

import "fmt"

func main() {
	// you can't put a funciton inside another function unless you do this...
	// we are assigning the results of an anonymous function to a variable,
	// then we can call that as a function
	greeting := func(name string) string {
		return fmt.Sprint("Hello ", name)
	}

	fmt.Println(greeting("Logan"))
}

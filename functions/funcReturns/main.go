package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
	hello := fmt.Sprint("hello world") // Sprint appears to print to string, but
	fmt.Println(hello)                 // still needs to be stored or used
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

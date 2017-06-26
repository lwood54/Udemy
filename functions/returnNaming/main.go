package main

import "fmt"

func main() {
	s := greet("Jane ", "Doe")
	fmt.Println(s, " is at address: ", &s)

}

// This shows that you identify a variable to be returned.
// Here we are telling the function that we will be returning
// the "s" value
func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	fmt.Println(s, " is at address: ", &s)
	return
}

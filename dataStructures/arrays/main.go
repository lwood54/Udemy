package main

import "fmt"

// Arrays have to have a predetermined length, and that can't be changed
// after they have been created.
func main() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = fmt.Sprintln(i, ": ", string(i))
	}
	x[42] = "another letter"
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}

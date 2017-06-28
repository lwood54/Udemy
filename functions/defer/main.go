package main

import "fmt"

func hello() {
	fmt.Println("hello ")
}

func world() {
	fmt.Println("world")
}

func myName() {
	fmt.Println("Logan")
}

// defer will defer that function until right before "main" exits
// it is often used to close a file that you had open
func main() {
	// no defer
	world()
	hello()

	// used defer
	// defer appears to work in reverse order, by that I mean that
	// it seems that the thing you FIRST use defer on while the code
	// is running, will be the LAST thing that will run...so it is the MOST DEFERRED
	// the last thing to be deferred will be the first of all the deferred statements
	// to run
	defer world()  // very last thing to run
	defer hello()  // 2nd to last thing to run
	defer myName() // 3rd to last thing to run

}

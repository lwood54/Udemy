package main

import "fmt"

func main() {
	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run.")
	}

	// initialization statement:
	b := true

	if food := "Chocolate"; b { // you can initialize inside the boolean if statement
		fmt.Println(food) // this helps keep the scope extremely narrow
		// can't access food outside this if statement
	}
}

package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0:  "Good morning",
		01: "Bonjour!",
		02: "Buenos dias!",
		03: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	// delete(myGreeting, 2)

	// NOTE: we can put this whole long expression in the conditional statement.
	// Notice that we could put val, exists := myGreeting[2] above, but then it would
	// be in the main block.
	// Instead, the scope of that statement is now confined to just the initial
	// if block statement.
	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)
}

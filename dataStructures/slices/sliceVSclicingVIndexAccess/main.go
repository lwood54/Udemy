// slice vs slicing vs index access
package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)       // [a b c g m z]
	fmt.Println(mySlice[2:4])  // [c g] --> start index to < limit
	fmt.Println(mySlice[2])    // [c]
	fmt.Println("myString"[2]) // 83 prints unicode

	// does the original slice change after you have sliced?
	oldSlice := []int{1, 2, 3, 4}
	fmt.Println(oldSlice)
	fmt.Println(len(oldSlice))

	newSlice := oldSlice[1:3]
	fmt.Println("newSlice is: ", newSlice)
	fmt.Println("length of newSlice is: ", len(newSlice))
	fmt.Println("oldSlice is now: ", oldSlice)
	fmt.Println("length of oldSlice is now: ", len(oldSlice))
	// slicing from the slice does not appear to change the original slice
}

package main

import "fmt"

func main() {
	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)

	// to delete, you basically just append to the slice, but
	// leave out the positions you don't want to carry over
	// So below, we append up to index 2, but skip over it,
	// start back at index 3 and continue to the end
	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice)
}

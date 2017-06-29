package main

import "fmt"

// We can make a slice to fit efficiently at the beginning.
// As the slice grows, the underlying array will be fileld in.
// The length of the slice is however many values there are in the slice.
// The capacity is is many possible values the underlying array can hold.

// If the slice grows to exceed the capacity, then a new array with double
// the capacity of the previous array will become the underlying array and
// the old array will be tossed (garbage collected I assume)
// Each time the slice fills up and will exceed the underlying array capacity,
// a new array will be built and will be double the length of the previous underlying array.

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("----------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Capacity: ", cap(mySlice), "Value: ", mySlice[i])
	}
}

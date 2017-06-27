package main

import "fmt"

// this function takes a slice of ints and a function that takes an int and returns a bool
// this function itself returns a slice of ints
func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}               // creating the variable xs that is a blank slice of ints
	for _, n := range numbers { // iterate through the slice of ints
		if callback(n) { // pulling in the callback function, if the value that is sent through callback
			xs = append(xs, n) // is true, then the value of xs is now the slice of ints with the value of n added to it
		}
	}
	return xs // returning the slice of ints with whatever values came back true
}

func main() {
	// in main, we created the value xs and set it equal to the result of the
	// filter function that passes the slice of ints: 1,2,3,4, and the anonymous
	// function that checks whether the n that is passed through is greater than 1
	// if n is greater than 1, the number returns true and is sent to the filter to
	// be added to the slice of ints
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	// we now print out the slice of ints, which should be values greater than 1
	fmt.Println(xs) // [2 3 4]

	// just to test, I'm going to have the same result as above, use the filter(),
	// but I also want to pass a function that is already made that will determine
	// if n > 1

	xy := filter([]int{1, 2, 3, 4}, compToOne)
	fmt.Println(xy)
}

func compToOne(num int) bool {
	return num > 1
}

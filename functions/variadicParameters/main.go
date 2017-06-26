package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(nums ...float64) float64 {
	var total float64
	for _, val := range nums { // range returns an index, value so we are putting the _ (aka blank identifier)
		total += val // here because we aren't interested which index each one is, we are just
	} // just using the value

	fmt.Println(nums)                 // this is a "slice" which is basically a list, it is of Type float64
	fmt.Printf("%T \n", nums)         // shows what Type nums is
	return total / float64(len(nums)) // float64() converts to float in order to divide total
} // which is float64 type

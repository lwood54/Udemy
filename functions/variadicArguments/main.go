package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data...) // when you pass an argument with ..., it tells the
	fmt.Println(n)        // func that you want to open up that list and pass each item
	// as an argument

}

func average(nums ...float64) float64 {
	var total float64
	for _, val := range nums {
		total += val
	}
	return total / float64(len(nums))
}

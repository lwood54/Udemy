// 3. Write a function with one variadic parameter that finds the greatest number
//     in a list of numbers.

package main

import "fmt"

func main() {
	fmt.Println(greatestNum([]int{4, 5, 66, 78, 2, 105, 8})) // should return 105
	fmt.Println(greatestNum([]int{540, 35, 25, 349, 9}))     // should return 540
}

func greatestNum(numList []int) int {
	var greatest int
	for _, val := range numList {
		if greatest < val {
			greatest = val
		}
	}
	return greatest
}

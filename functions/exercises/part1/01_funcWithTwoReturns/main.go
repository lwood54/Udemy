// 1. Write a function which takes an integer. The function will have two returns.
//     The first return should be the argument divided by 2.
//     The second return should be a bool that lets us know whether or not the
//     argument was even. For example:
//       a. half(1) returns (0, false)
//       b. half(2) returns (1, true)

package main

import "fmt"

func main() {
	var halfVal int
	var isEven bool
	halfVal, isEven = half(1)
	fmt.Println(halfVal, isEven) // should be 0 false

	halfVal, isEven = half(2)
	fmt.Println(halfVal, isEven) // should be 1 true

	halfVal, isEven = half(100)
	fmt.Println(halfVal, isEven) // should be 50 true

}

func half(num int) (int, bool) {
	var isEven bool
	var newNum int
	if num%2 == 0 {
		isEven = true
	}
	newNum = num / 2
	return newNum, isEven
}

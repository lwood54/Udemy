// 2. Modify the previous program to use a func expression.
//   NOTE: Remember an expression has a value and is generally written horizontally.
//         A statement is written vertically and can contain multiple expressions.
//         (Kind of like a sentence vs a paragraph.)

package main

import "fmt"

func main() {
	half := func(num int) (int, bool) {
		var isEven bool
		var newNum int
		if num%2 == 0 {
			isEven = true
		}
		newNum = num / 2
		return newNum, isEven
	}

	var halfVal int
	var isEven bool
	halfVal, isEven = half(1)
	fmt.Println(halfVal, isEven) // should be 0 false

	halfVal, isEven = half(2)
	fmt.Println(halfVal, isEven) // should be 1 true

	halfVal, isEven = half(100)
	fmt.Println(halfVal, isEven) // should be 50 true

}

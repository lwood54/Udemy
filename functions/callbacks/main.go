package main

import "fmt"

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	fmt.Println("It worked.")
}

// a func called visit() has to parameters
// the first parameter will be a basically a list of type slice int
// the second parameter is the variable callback that takes a function of type func that takes an int that has no return
// the visit() func then loops through the range in the slice and sends it to the callback() func
// so it's passing the whole func(n int) { fmt.Println(n)} to the visit().
// the visit() func takes the list, loops through, has an n, that is plugged in
// to the variable that represents the anonymous func that Println's n
func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

// callback: passing a func as an argument

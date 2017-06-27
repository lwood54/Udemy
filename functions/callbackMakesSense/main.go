package main

import "fmt"

func main() {
	iterate([]int{2, 4, 6, 8}, func(n int) {
		fmt.Println(n)
	})

	iterate([]int{2, 4, 6, 8}, func(num int) {
		fmt.Println(num + 5)
	})

	// IMPORTANT: the key for passing a named function instead of just
	// an anonymous function is to make sure NOT to add the () at the end
	// of the function name. This way it passes the value of the function
	// instead of trying to run it inside the other function.
	iterate([]int{2, 4, 6, 8}, addTen)

	phrase := hello(logan())
	fmt.Println(phrase)

} // end of func main()

// The func below didn't work as a callback. I'm wondering if you can only use
// anonymous functions
func addTen(num int) {
	fmt.Println("It worked!!!")
	fmt.Println(num + 10)
}

func logan() string {
	return "Logan"
}

func hello(name string) string {
	return fmt.Sprint("Hello ", name)
}

func iterate(numList []int, useThisFunc func(int)) {
	for _, val := range numList {
		useThisFunc(val)
	}
}
